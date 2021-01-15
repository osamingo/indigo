package base58_test

import (
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/osamingo/indigo/base58"
)

func TestStdSource(t *testing.T) {
	const expect = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	if base58.StdSource() != expect {
		t.Error("should be", expect)
	}
}

func TestMustNewEncoder(t *testing.T) {

	enc := base58.MustNewEncoder("rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyz")
	if enc == nil {
		t.Error("should not be nil")
	}

	func() {
		defer func() {
			recover()
		}()
		base58.MustNewEncoder("")
		t.Error("should be panic")
	}()

	func() {
		defer func() {
			recover()
		}()
		base58.MustNewEncoder("test")
		t.Error("should be panic")
	}()
}

func TestNewEncoder(t *testing.T) {

	enc, err := base58.NewEncoder("rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyz")
	if err != nil {
		t.Error("should be nil")
	}
	if enc == nil {
		t.Error("should not be nil")
	}

	_, err = base58.NewEncoder("")
	if err == nil {
		t.Error("should not be nil")
	}

	_, err = base58.NewEncoder("test")
	if err == nil {
		t.Error("should not be nil")
	}
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
	if id != "1" {
		t.Error("should be", "1")
	}

	for k, v := range bc {
		if enc.Encode(k) != v {
			t.Error("should be", v)
		}
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
	if err == nil {
		t.Error("should not be nil")
	}

	_, err = enc.Decode("0")
	if err == nil {
		t.Error("should not be nil")
	}

	for k, v := range bc {
		r, err := enc.Decode(v)
		if err != nil {
			t.Error("should be nil")
		}
		if r != k {
			t.Error("should be", k)
		}
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
