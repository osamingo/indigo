package indigo_test

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"sync"
	"testing"
	"time"

	"github.com/osamingo/indigo"
	"github.com/osamingo/indigo/base58"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	g := indigo.New(
		base58.MustNewEncoder(base58.StdSource()),
		indigo.StartTime(time.Unix(1257894000, 0)),
		indigo.MachineID(func() (uint16, error) { return math.MaxUint16, nil }),
		indigo.CheckMachineID(nil),
	)
	require.NotNil(t, g)
}

func TestGenerator_NextID(t *testing.T) {

	g := indigo.New(
		nil,
		indigo.StartTime(time.Unix(1257894000, 0)),
		indigo.MachineID(func() (uint16, error) { return math.MaxUint16, nil }),
		indigo.CheckMachineID(nil),
	)

	id1, err := g.NextID()
	require.NoError(t, err)
	assert.NotEmpty(t, id1)

	id2, err := g.NextID()
	require.NoError(t, err)
	assert.NotEmpty(t, id2)
	assert.NotEqual(t, id1, id2)
}

func TestGenerator_Decompose(t *testing.T) {

	g := indigo.New(
		nil,
		indigo.StartTime(time.Unix(1257894000, 0)),
		indigo.MachineID(func() (uint16, error) { return math.MaxUint16, nil }),
		indigo.CheckMachineID(nil),
	)

	m, err := g.Decompose("KGuFE14P")
	require.NoError(t, err)
	require.NotEmpty(t, m)
	assert.NotEmpty(t, m["id"])

	_, err = g.Decompose("")
	require.Error(t, err)
}

func TestGenerator_NextID_Race(t *testing.T) {

	g := indigo.New(
		nil,
		indigo.StartTime(time.Unix(1257894000, 0)),
		indigo.MachineID(func() (uint16, error) { return math.MaxUint16, nil }),
	)

	gs := 2048

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			defer wg.Done()
			id, err := g.NextID()
			require.NoError(t, err)
			require.NotEmpty(t, id)
		}()
	}

	wg.Wait()
}

func TestGenerator_NextID_SortIDs(t *testing.T) {

	th := 10
	ids := make([]string, 0, 100)

	m := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(th)

	for i := 0; i < th; i++ {
		go func(mm uint16) {
			defer wg.Done()

			g := indigo.New(
				nil,
				indigo.StartTime(time.Unix(1257894000, 0)),
				indigo.MachineID(func() (uint16, error) { return mm, nil }),
			)

			r := rand.New(rand.NewSource(time.Now().UnixNano()))

			s := make([]string, 0, th)
			for j := 0; j < th; j++ {
				time.Sleep(10*time.Millisecond + time.Duration(r.Intn(1e9)))
				id, err := g.NextID()
				require.NoError(t, err)
				s = append(s, id)
			}

			m.Lock()
			ids = append(ids, s...)
			m.Unlock()
		}(uint16(i + 1))
	}

	wg.Wait()

	old := make([]string, 100)
	copy(old, ids)
	require.Equal(t, old, ids)

	sort.Strings(ids)
	require.NotEqual(t, old, ids)

	g := indigo.New(
		nil,
		indigo.StartTime(time.Unix(1257894000, 0)),
		indigo.MachineID(func() (uint16, error) { return math.MaxUint16, nil }),
	)

	var prev uint64
	for i := range ids {
		m, err := g.Decompose(ids[i])
		require.NoError(t, err)
		require.True(t, prev <= m["time"])
		prev = m["time"]
	}
}

func BenchmarkGenerator_NextID(b *testing.B) {

	g := indigo.New(
		nil,
		indigo.StartTime(time.Now()),
		indigo.MachineID(func() (uint16, error) { return math.MaxUint16, nil }),
	)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		id, err := g.NextID()
		if err != nil {
			b.Fatal(err)
		}
		if id == "" {
			b.Fatal("generate id is empty")
		}
	}
}

func ExampleGenerator_NextID() {
	const machineID = 65535
	g := indigo.New(
		nil,
		indigo.StartTime(time.Now()),
		indigo.MachineID(func() (uint16, error) { return machineID, nil }),
	)
	id, err := g.NextID()
	if err != nil {
		panic(err)
	}
	fmt.Println(id)

	m, err := g.Decompose(id)
	if err != nil {
		panic(err)
	}
	fmt.Println(m["machine-id"])
	// output:
	// 2VKmG
	// 65535
}
