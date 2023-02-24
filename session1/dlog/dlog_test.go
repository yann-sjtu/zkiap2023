package dlog

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVerify(t *testing.T) {
	const p = 31
	const g = 3
	const x = 17
	//const y = 22
	y, pf := dLogProof(x, g, p)
	require.Equal(t, 22, y)
	require.True(t, verify(y, g, p, pf))
}

func TestExp(t *testing.T) {
	largeN := 1 << 62
	require.Equal(t, 16, exp(2, 4, largeN))
	require.Equal(t, 1024, exp(2, 10, largeN))
	require.Equal(t, 1024*32, exp(2, 15, largeN))
	require.Equal(t, 1024*1024, exp(2, 20, largeN))
	require.Equal(t, 3125, exp(5, 5, largeN))
	require.Equal(t, 3125*25, exp(5, 7, largeN))

	require.Equal(t, 32, exp(2, 15, 1023))
	require.Equal(t, 169, exp(2, 20, 1011))
}
