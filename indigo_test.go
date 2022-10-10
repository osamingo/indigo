package indigo_test

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"sort"
	"sync"
	"testing"
	"time"

	"github.com/osamingo/base58"
	"github.com/osamingo/indigo"
)

func TestNew(t *testing.T) {
	t.Parallel()

	g := indigo.New(
		base58.MustNewEncoder(base58.StandardSource),
		indigo.StartTime(time.Unix(1257894000, 0)),
		indigo.MachineID(func() (uint16, error) { return math.MaxUint16, nil }),
		indigo.CheckMachineID(nil),
	)
	if g == nil {
		t.Error("should not be nil")
	}
}

func TestGenerator_NextID(t *testing.T) {
	t.Parallel()

	g := indigo.New(
		nil,
		indigo.StartTime(time.Unix(1257894000, 0)),
		indigo.MachineID(func() (uint16, error) { return math.MaxUint16, nil }),
		indigo.CheckMachineID(nil),
	)

	id1, err := g.NextID()
	if err != nil {
		t.Error("should be nil")
	}

	if id1 == "" {
		t.Error("should not be empty")
	}

	id2, err := g.NextID()
	if err != nil {
		t.Error("should be nil")
	}

	if id2 == "" {
		t.Error("should not be empty")
	}

	if id1 == id2 {
		t.Error("should not be equal")
	}

	g = indigo.New(
		nil,
		indigo.StartTime(time.Unix(math.MinInt64, math.MinInt64)),
		indigo.MachineID(func() (uint16, error) { return math.MaxUint16, nil }),
	)

	if _, err := g.NextID(); err == nil {
		t.Error("should not be nil")
	}
}

func TestGenerator_Decompose(t *testing.T) {
	t.Parallel()

	g := indigo.New(
		nil,
		indigo.StartTime(time.Unix(1257894000, 0)),
		indigo.MachineID(func() (uint16, error) { return math.MaxUint16, nil }),
		indigo.CheckMachineID(nil),
	)

	m, err := g.Decompose("KGuFE14P")
	if err != nil {
		t.Error("should be nil")
	}

	if len(m) == 0 {
		t.Error("should not be empty")
	}

	if _, ok := m["id"]; !ok {
		t.Error("should not be empty")
	}

	_, err = g.Decompose("")
	if err == nil {
		t.Error("should not be nil")
	}
}

func TestGenerator_NextID_Race(t *testing.T) {
	t.Parallel()

	g := indigo.New(
		nil,
		indigo.StartTime(time.Unix(1257894000, 0)),
		indigo.MachineID(func() (uint16, error) { return math.MaxUint16, nil }),
	)

	gs := 2048
	wg := sync.WaitGroup{}

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			defer wg.Done()

			id, err := g.NextID()
			if err != nil {
				t.Error("should be nil")
			}

			if id == "" {
				t.Error("should not be empty")
			}
		}()
	}

	wg.Wait()
}

func TestGenerator_NextID_SortIDs(t *testing.T) {
	t.Parallel()

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
				if err != nil {
					t.Error("should be nil")
				}

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

	if !reflect.DeepEqual(ids, old) {
		t.Error("should be equal")
	}

	sort.Strings(ids)

	if reflect.DeepEqual(ids, old) {
		t.Error("should not be equal")
	}

	g := indigo.New(
		nil,
		indigo.StartTime(time.Unix(1257894000, 0)),
		indigo.MachineID(func() (uint16, error) { return math.MaxUint16, nil }),
	)

	var prev uint64

	for i := range ids {
		m, err := g.Decompose(ids[i])
		if err != nil {
			t.Error("should be nil")
		}

		if !(prev <= m["time"]) {
			t.Error("should be true")
		}

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

func ExampleGenerator_NextID() { //nolint: nosnakecase
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
