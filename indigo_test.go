package indigo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var mid = func() (uint16, error) { return 0, nil }

func TestIndigo(t *testing.T) {

	require.Nil(t, s)

	id, err := NextID()
	require.Error(t, err)

	require.NotPanics(t, func() { New(time.Now(), mid, nil) })

	id, err = NextID()
	require.NoError(t, err)
	require.NotEmpty(t, id)
}

func TestDecompose(t *testing.T) {

	m, err := Decompose("KGuFE14P")
	require.NoError(t, err)
	require.NotEmpty(t, m)
	assert.NotEmpty(t, m["id"])

	_, err = Decompose("")
	require.Error(t, err)
}

func TestRaceNextID(t *testing.T) {
	for i := 0; i < 2048; i++ {
		go func() {
			id, err := NextID()
			require.NoError(t, err)
			require.NotEmpty(t, id)
		}()
	}
}

func BenchmarkNextID(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NextID()
	}
}
