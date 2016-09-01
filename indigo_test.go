package indigo

import (
	"testing"
	"time"

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

	m, err := Decompose(id)
	require.NoError(t, err)
	require.NotEmpty(t, m)

	_, err = Decompose("")
	require.Error(t, err)
}

func TestRaceNextID(t *testing.T) {

	New(time.Now(), mid, nil)

	for i := 0; i < 2048; i++ {
		go func() {
			id, err := NextID()
			require.NoError(t, err)
			require.NotEmpty(t, id)
		}()
	}
}

func BenchmarkNextID(b *testing.B) {

	New(time.Now(), mid, nil)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NextID()
	}
}
