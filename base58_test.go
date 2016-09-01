package indigo

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var tc = map[uint64]string{
	0:          "1",
	32:         "y",
	64:         "27",
	512:        "9Q",
	1024:       "iE",
	2048:       "Bj",
	4096:       "2dC",
	8192:       "3rf",
	16384:      "5Su",
	32768:      "aJY",
	65536:      "ktW",
	131072:     "EXS",
	262144:     "2kVJ",
	524288:     "3FRs",
	1048576:    "6nGU",
	16777216:   "2tZhm",
	134217728:  "cRUfL",
	1073741824: "2CTd35",
	// uint64: MAX_VALUE
	18446744073709551615: "JPwcyDCgEup",
}

func TestEncodeBase58(t *testing.T) {
	for k, v := range tc {
		assert.Equal(t, v, EncodeBase58(k))
	}
}

func TestDecodeBase58(t *testing.T) {
	for k, v := range tc {
		r, err := DecodeBase58(v)
		require.NoError(t, err)
		assert.Equal(t, k, r)
	}

	_, err := DecodeBase58("0")
	require.Error(t, err)
}

func BenchmarkEncodeBase58(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := range tc {
			EncodeBase58(i)
		}
	}
}

func BenchmarkDecodeBase58(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := range tc {
			DecodeBase58(tc[i])
		}
	}
}
