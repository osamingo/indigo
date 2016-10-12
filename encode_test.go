package indigo

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerator_Encode(t *testing.T) {

	g := New(Settings{
		StartTime: time.Now(),
		MachineID: mid,
	})

	for k, v := range tc {
		assert.Equal(t, v, g.Encode(k))
	}
}

func BenchmarkGenerator_Encode(b *testing.B) {

	g := New(Settings{
		StartTime: time.Now(),
		MachineID: mid,
	})

	s := rand.New(rand.NewSource(time.Now().UnixNano()))

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.Encode(uint64(s.Int63()))
	}
}
