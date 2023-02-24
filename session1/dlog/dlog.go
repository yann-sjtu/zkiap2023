package dlog

import (
	"math/rand"
)

const rounds = 100

func dLogProof(x, g, p int) (y int, pf [rounds * 2]int) {
	y = exp(g, x, p)

	rs := genr(p)
	var hs [rounds]int
	for i := 0; i < rounds; i++ {
		hs[i] = exp(g, rs[i], p)
	}

	copy(pf[:rounds], hs[:])

	bs := genb(hs)
	for i := 0; i < rounds; i++ {
		if bs[i] == 0 {
			pf[i+rounds] = rs[i] % (p - 1)
		} else {
			pf[i+rounds] = (rs[i] + x) % (p - 1)
		}
	}

	return
}

func verify(y, g, p int, pf [rounds * 2]int) bool {
	var hs [rounds]int
	copy(hs[:], pf[:rounds])
	bs := genb(hs)
	for i := 0; i < rounds; i++ {
		if bs[i] == 0 && exp(g, pf[i+rounds], p) != hs[i] {
			return false
		} else if bs[i] == 1 && exp(g, pf[i+rounds], p) != hs[i]*y%p {
			return false
		}
	}
	return true
}

func genr(p int) [rounds]int {
	var rs [rounds]int
	for i := 0; i < rounds; i++ {
		rs[i] = rand.Intn(p - 1)
	}
	return rs
}

// generate r -> calculate h -> generate b
func genb(hs [rounds]int) [rounds]int {
	var bs [rounds]int
	for i := 0; i < rounds; i++ {
		bs[i] = hs[i] >> 2 & 1
	}
	return bs

}
