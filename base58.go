package indigo

import (
	"encoding/binary"
	"errors"
	"sort"
	"strings"
)

const base58 = 58

var (
	decodeMap  = make([]int64, 256)
	characters = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
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

// SetBase58Characters changes characters of Base58.
func SetBase58Characters(chars string, sorting bool) error {
	if len(chars) != base58 {
		return errors.New("indigo: characters must be 58 length")
	}

	if sorting {
		s := strings.Split(chars, "")
		sort.Strings(s)
		chars = strings.Join(s, "")
	}

	characters = chars
	defineDecodeMap()
	return nil
}

// EncodeBase58 returns encoded string by Base58.
func EncodeBase58(u uint64) string {

	if u == 0 {
		return characters[:1]
	}

	d := make([]byte, 0, binary.MaxVarintLen64)
	for u > 0 {
		d = append(d, characters[u%base58])
		u /= base58
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
