package indigo

import (
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var tc = map[uint64]string{
	0:              "1",
	57:             "z",
	math.MaxUint8:  "5Q",
	math.MaxUint16: "LUv",
	math.MaxUint32: "7YXq9G",
	math.MaxUint64: "jpXCZedGfVQ",
}

func TestSetBase58Characters(t *testing.T) {

	orig := characters

	err := SetBase58Characters("", false)
	require.Error(t, err)

	err = SetBase58Characters("123", false)
	require.Error(t, err)

	ripple := "rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyz"

	err = SetBase58Characters(ripple, false)
	require.NoError(t, err)
	assert.Equal(t, "rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyz", characters)

	err = SetBase58Characters(ripple, true)
	require.NoError(t, err)
	assert.Equal(t, "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz", characters)

	err = SetBase58Characters(orig, true)
	require.NoError(t, err)
	assert.Equal(t, orig, characters)
}

func TestEncodeBase58(t *testing.T) {
	for k, v := range tc {
		assert.Equal(t, v, EncodeBase58(k))
	}
}

func TestDecodeBase58(t *testing.T) {
	for k, v := range tc {
		r, err := DecodeBase58(v)
		require.NoError(t, err)
		assert.Equal(t, k, r)
	}

	_, err := DecodeBase58("0")
	require.Error(t, err)
}

func BenchmarkEncodeBase58(b *testing.B) {
	s := rand.NewSource(time.Now().UnixNano())
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EncodeBase58(uint64(s.Int63()))
	}
}

func BenchmarkDecodeBase58(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := range tc {
			DecodeBase58(tc[i])
		}
	}
}
