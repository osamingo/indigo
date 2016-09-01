package indigo

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var tc = map[uint64]string{
	0:           "1",
	32:          "y",
	64:          "27",
	512:         "9Q",
	1024:        "iE",
	16777216:    "2tZhm",
	68719476736: "2NGvhhq",
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
