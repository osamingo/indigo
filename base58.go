package indigo

import (
	"encoding/binary"
	"errors"
)

const base58 = 58

var (
	decodeMap  = make([]int64, 256)
	characters = "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"
)

func init() {
	defineDecodeMap()
}

func defineDecodeMap() {
	for i := range decodeMap {
		decodeMap[i] = -1
	}
	for i := range characters {
		decodeMap[characters[i]] = int64(i)
	}
}

// EncodeBase58 returns encoded string by Base58.
func EncodeBase58(u uint64) string {

	if u == 0 {
		return characters[:1]
	}

	d := make([]byte, 0, binary.MaxVarintLen64)
	for u > 0 {
		d, u = append(d, characters[u%base58]), u/base58
	}

	for i, j := 0, len(d)-1; i < j; i, j = i+1, j-1 {
		d[i], d[j] = d[j], d[i]
	}

	return string(d)
}

// DecodeBase58 returns decoded unsigned int64 by Base58.
func DecodeBase58(s string) (uint64, error) {

	if len(s) == 0 {
		return 0, errors.New("indigo: source should not be empty")
	}

	n := uint64(0)
	for i := range s {
		u := decodeMap[s[i]]
		if u < 0 {
			return 0, errors.New("indigo: invalid character = " + string(s[i]))
		}
		n = n*base58 + uint64(u)
	}
	return n, nil
}
