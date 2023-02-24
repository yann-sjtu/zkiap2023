package signature

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVerify(t *testing.T) {
	message := "test message"
	privateKey := random(message).String()

	pf := sign(message, privateKey)

	require.True(t, verify(message, pf))

	pf2 := pf.deepCopy()
	pf2.md5[0]++
	require.False(t, verify(message, pf2))

	pf3 := pf.deepCopy()
	pf3.gx[0]++
	require.False(t, verify(message, pf3))

	pf4 := pf.deepCopy()
	pf4.mx[0]++
	require.False(t, verify(message, pf4))

	pf5 := pf.deepCopy()
	pf5.mr[0]++
	require.False(t, verify(message, pf5))

	pf6 := pf.deepCopy()
	pf6.gr[0]++
	require.False(t, verify(message, pf6))

	pf7 := pf.deepCopy()
	pf7.s[0]++
	require.False(t, verify(message, pf7))
}

func (pf proof) deepCopy() proof {
	pf2 := proof{
		md5: pf.md5,
		gx:  make([]byte, len(pf.gx)),
		mx:  make([]byte, len(pf.mx)),
		mr:  make([]byte, len(pf.mr)),
		gr:  make([]byte, len(pf.gr)),
		s:   make([]byte, len(pf.s)),
	}
	copy(pf2.gx, pf.gx)
	copy(pf2.mx, pf.mx)
	copy(pf2.mr, pf.mr)
	copy(pf2.gr, pf.gr)
	copy(pf2.s, pf.s)
	return pf2
}
