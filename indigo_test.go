package indigo

import (
	"math/rand"
	"sort"
	"sync"
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

	gs := 2048

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			defer wg.Done()
			id, err := NextID()
			require.NoError(t, err)
			require.NotEmpty(t, id)
		}()
	}

	wg.Wait()
}

func TestOrderedIDs(t *testing.T) {

	ids := make([]string, 10)

	var err error
	for i := range ids {
		time.Sleep(10 * time.Millisecond)
		ids[i], err = NextID()
		require.NoError(t, err)
	}

	for i := range ids {
		j := rand.Intn(i + 1)
		ids[i], ids[j] = ids[j], ids[i]
	}

	old := make([]string, 10)

	copy(old, ids)
	require.Equal(t, old, ids)

	sort.Strings(ids)
	require.NotEqual(t, old, ids)

	var prev uint64
	for i := range ids {
		m, err := Decompose(ids[i])
		require.NoError(t, err)
		require.True(t, prev < m["time"])
		prev = m["time"]
	}
}

func BenchmarkNextID(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NextID()
	}
}
