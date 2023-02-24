package signature

import (
	"crypto/md5"
	"crypto/sha256"
	"math/big"
)

var (
	p, _ = new(big.Int).SetString("273389558745553615023177755634264971227", 10)
	g, _ = new(big.Int).SetString("191981998178538467192271372964660528157", 10)
)

type proof struct {
	md5 [md5.Size]byte
	gx  []byte
	mx  []byte
	mr  []byte
	gr  []byte
	s   []byte
}

func sign(message, privateKey string) proof {
	x, ok := new(big.Int).SetString(privateKey, 10)
	if !ok {
		panic("invalid private key")
	}

	gx := new(big.Int).Exp(g, x, p)

	md5m := md5.Sum([]byte(message))
	m := new(big.Int).SetBytes(md5m[:])
	r := random(message + privateKey)

	mx := new(big.Int).Exp(m, x, p)
	mr := new(big.Int).Exp(m, r, p)
	gr := new(big.Int).Exp(g, r, p)
	c := genC(mx, mr, gr)
	s := new(big.Int).Mul(c, x)
	s.Add(s, r)

	return proof{
		md5: md5m,
		gx:  gx.Bytes(),
		mx:  mx.Bytes(),
		mr:  mr.Bytes(),
		gr:  gr.Bytes(),
		s:   s.Bytes(),
	}
}

func verify(message string, proof proof) bool {
	if md5.Sum([]byte(message)) != proof.md5 {
		return false
	}

	m := new(big.Int).SetBytes(proof.md5[:])
	gx := new(big.Int).SetBytes(proof.gx)
	mx := new(big.Int).SetBytes(proof.mx)
	mr := new(big.Int).SetBytes(proof.mr)
	gr := new(big.Int).SetBytes(proof.gr)
	s := new(big.Int).SetBytes(proof.s)

	c := genC(mx, mr, gr)
	gs := new(big.Int).Exp(gx, c, p)
	gs = gs.Mul(gs, gr)
	gs = gs.Mod(gs, p)
	if new(big.Int).Exp(g, s, p).Cmp(gs) != 0 {
		return false
	}

	ms := new(big.Int).Exp(mx, c, p)
	ms = ms.Mul(ms, mr)
	ms = ms.Mod(ms, p)
	if new(big.Int).Exp(m, s, p).Cmp(ms) != 0 {
		return false
	}

	return true
}

func genC(mx, mr, gr *big.Int) *big.Int {
	var seed []byte
	seed = append(seed, mx.Bytes()...)
	seed = append(seed, mr.Bytes()...)
	seed = append(seed, gr.Bytes()...)
	hash := sha256.Sum256(seed)
	return new(big.Int).SetBytes(hash[:])
}

func random(data string) *big.Int {
	hash := sha256.Sum256([]byte(data))
	n := new(big.Int).SetBytes(hash[:])
	return n.Mod(n, p)
}
