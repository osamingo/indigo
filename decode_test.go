package indigo

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerator_Decode(t *testing.T) {

	g := New(Settings{
		StartTime: time.Now(),
		MachineID: mid,
	})

	_, err := g.Decode("0")
	require.Error(t, err)

	for k, v := range tc {
		r, err := g.Decode(v)
		require.NoError(t, err)
		assert.Equal(t, k, r)
	}
}

func BenchmarkGenerator_Decode(b *testing.B) {

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
		g.Decode(vs[s.Intn(l)])
	}
}
