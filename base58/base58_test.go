package base58_test

import (
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/osamingo/indigo/base58"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStdSource(t *testing.T) {
	require.Equal(t, base58.StdSource(), "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")
}

func TestMustNewEncoder(t *testing.T) {

	var enc *base58.Encoder
	require.NotPanics(t, func() {
		enc = base58.MustNewEncoder("rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyz")
	})
	require.NotNil(t, enc)

	require.Panics(t, func() {
		base58.MustNewEncoder("")
	})

	require.Panics(t, func() {
		base58.MustNewEncoder("test")
	})
}

func TestNewEncoder(t *testing.T) {

	enc, err := base58.NewEncoder("rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyz")
	require.NoError(t, err)
	require.NotNil(t, enc)

	_, err = base58.NewEncoder("")
	require.Error(t, err)

	_, err = base58.NewEncoder("test")
	require.Error(t, err)
}

func TestEncoder_Encode(t *testing.T) {

	bc := map[uint64]string{
		0:              "1",
		57:             "z",
		math.MaxUint8:  "5Q",
		math.MaxUint16: "LUv",
		math.MaxUint32: "7YXq9G",
		math.MaxUint64: "jpXCZedGfVQ",
	}

	enc := base58.MustNewEncoder(base58.StdSource())
	id := enc.Encode(0)
	assert.Equal(t, "1", id)

	for k, v := range bc {
		assert.Equal(t, v, enc.Encode(k))
	}
}

func TestEncoder_Decode(t *testing.T) {

	bc := map[uint64]string{
		0:              "1",
		57:             "z",
		math.MaxUint8:  "5Q",
		math.MaxUint16: "LUv",
		math.MaxUint32: "7YXq9G",
		math.MaxUint64: "jpXCZedGfVQ",
	}

	enc := base58.MustNewEncoder(base58.StdSource())
	_, err := enc.Decode("")
	require.Error(t, err)

	_, err = enc.Decode("0")
	require.Error(t, err)

	for k, v := range bc {
		r, err := enc.Decode(v)
		require.NoError(t, err)
		assert.Equal(t, k, r)
	}
}

func BenchmarkEncoder_Encode(b *testing.B) {

	s := rand.New(rand.NewSource(time.Now().UnixNano()))
	enc := base58.MustNewEncoder(base58.StdSource())

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		enc.Encode(uint64(s.Int63()))
	}
}

func BenchmarkEncoder_Decode(b *testing.B) {

	bc := map[uint64]string{
		0:              "1",
		57:             "z",
		math.MaxUint8:  "5Q",
		math.MaxUint16: "LUv",
		math.MaxUint32: "7YXq9G",
		math.MaxUint64: "jpXCZedGfVQ",
	}

	l := len(bc)
	s := rand.New(rand.NewSource(time.Now().UnixNano()))
	enc := base58.MustNewEncoder(base58.StdSource())

	vs := make([]string, 0, l)
	for k := range bc {
		vs = append(vs, bc[k])
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := enc.Decode(vs[s.Intn(l)])
		if err != nil {
			b.Fatal(err)
		}
	}
}
