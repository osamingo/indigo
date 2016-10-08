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

func TestGenerator_EncodeBase58(t *testing.T) {

	g := New(Settings{
		StartTime: time.Now(),
		MachineID: mid,
	})

	for k, v := range tc {
		assert.Equal(t, v, g.EncodeBase58(k))
	}
}

func TestGenerator_DecodeBase58(t *testing.T) {

	g := New(Settings{
		StartTime: time.Now(),
		MachineID: mid,
	})

	_, err := g.DecodeBase58("0")
	require.Error(t, err)

	for k, v := range tc {
		r, err := g.DecodeBase58(v)
		require.NoError(t, err)
		assert.Equal(t, k, r)
	}
}

func BenchmarkGenerator_EncodeBase58(b *testing.B) {

	g := New(Settings{
		StartTime: time.Now(),
		MachineID: mid,
	})

	s := rand.New(rand.NewSource(time.Now().UnixNano()))

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.EncodeBase58(uint64(s.Int63()))
	}
}

func BenchmarkGenerator_DecodeBase58(b *testing.B) {

	g := New(Settings{
		StartTime: time.Now(),
		MachineID: mid,
	})

	l := len(tc)
	s := rand.New(rand.NewSource(time.Now().UnixNano()))

	vs := make([]string, 0, l)
	for k := range tc {
		vs = append(vs, tc[k])
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.DecodeBase58(vs[s.Intn(l)])
	}
}
