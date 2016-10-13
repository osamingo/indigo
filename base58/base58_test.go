package base58

import (
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var bc = map[uint64]string{
	0:              "1",
	57:             "z",
	math.MaxUint8:  "5Q",
	math.MaxUint16: "LUv",
	math.MaxUint32: "7YXq9G",
	math.MaxUint64: "jpXCZedGfVQ",
}

func TestNewMustEncoder(t *testing.T) {

	var enc *Encoder
	require.NotPanics(t, func() {
		enc = NewMustEncoder("rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyz")
	})
	require.NotNil(t, enc)

	require.Panics(t, func() {
		NewMustEncoder("")
	})

	require.Panics(t, func() {
		NewMustEncoder("test")
	})
}

func TestNewEncoder(t *testing.T) {

	enc, err := NewEncoder("rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyz")
	require.NoError(t, err)
	require.NotNil(t, enc)

	_, err = NewEncoder("")
	require.Error(t, err)

	_, err = NewEncoder("test")
	require.Error(t, err)
}

func TestEncoder_Encode(t *testing.T) {

	id := StdEncoding.Encode(0)
	assert.Equal(t, "1", id)

	for k, v := range bc {
		assert.Equal(t, v, StdEncoding.Encode(k))
	}
}

func TestEncoder_Decode(t *testing.T) {

	_, err := StdEncoding.Decode("")
	require.Error(t, err)

	_, err = StdEncoding.Decode("0")
	require.Error(t, err)

	for k, v := range bc {
		r, err := StdEncoding.Decode(v)
		require.NoError(t, err)
		assert.Equal(t, k, r)
	}
}

func BenchmarkEncoder_Encode(b *testing.B) {

	s := rand.New(rand.NewSource(time.Now().UnixNano()))

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StdEncoding.Encode(uint64(s.Int63()))
	}
}

func BenchmarkEncoder_Decode(b *testing.B) {

	l := len(bc)
	s := rand.New(rand.NewSource(time.Now().UnixNano()))

	vs := make([]string, 0, l)
	for k := range bc {
		vs = append(vs, bc[k])
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StdEncoding.Decode(vs[s.Intn(l)])
	}
}
