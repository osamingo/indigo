package indigo

import (
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var tc = map[uint64][]byte{
	0:              []byte("1"),
	57:             []byte("z"),
	math.MaxUint8:  []byte("5Q"),
	math.MaxUint16: []byte("LUv"),
	math.MaxUint32: []byte("7YXq9G"),
	math.MaxUint64: []byte("jpXCZedGfVQ"),
}

func TestSetBase58Characters(t *testing.T) {

	orig := string(characters)

	err := SetBase58Characters("", false)
	require.Error(t, err)

	err = SetBase58Characters("123", false)
	require.Error(t, err)

	ripple := "rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyz"

	err = SetBase58Characters(ripple, false)
	require.NoError(t, err)
	assert.Equal(t, "rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyz", string(characters))

	err = SetBase58Characters(ripple, true)
	require.NoError(t, err)
	assert.Equal(t, "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz", string(characters))

	err = SetBase58Characters(orig, true)
	require.NoError(t, err)
	assert.Equal(t, orig, string(characters))
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

	_, err := DecodeBase58([]byte("0"))
	require.Error(t, err)
}

func BenchmarkEncodeBase58(b *testing.B) {

	s := rand.New(rand.NewSource(time.Now().UnixNano()))

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EncodeBase58(uint64(s.Int63()))
	}
}

func BenchmarkDecodeBase58(b *testing.B) {

	l := len(tc)
	s := rand.New(rand.NewSource(time.Now().UnixNano()))

	vs := make([][]byte, 0, l)
	for k := range tc {
		vs = append(vs, tc[k])
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DecodeBase58(vs[s.Intn(l)])
	}
}
