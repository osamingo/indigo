package indigo

import (
	"encoding/binary"
	"errors"
	"sort"
	"strings"
)

const fiftyEight = 58

var (
	decodeMap  = make([]int64, 256)
	characters = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")
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
	if len(chars) != fiftyEight {
		return errors.New("indigo: characters must be 58 length")
	}

	if sorting {
		s := strings.Split(chars, "")
		sort.Strings(s)
		chars = strings.Join(s, "")
	}

	characters = []byte(chars)
	defineDecodeMap()
	return nil
}

// EncodeBase58 returns encoded byte slice by Base58.
func EncodeBase58(id uint64) []byte {

	if id == 0 {
		return characters[:1]
	}

	bin := make([]byte, 0, binary.MaxVarintLen64)
	for id > 0 {
		bin = append(bin, characters[id%fiftyEight])
		id /= fiftyEight
	}

	for i, j := 0, len(bin)-1; i < j; i, j = i+1, j-1 {
		bin[i], bin[j] = bin[j], bin[i]
	}

	return bin
}

// DecodeBase58 returns decoded unsigned int64 by Base58.
func DecodeBase58(id []byte) (uint64, error) {

	if len(id) == 0 {
		return 0, errors.New("indigo: source should not be empty")
	}

	var n uint64
	for i := range id {
		u := decodeMap[id[i]]
		if u < 0 {
			return 0, errors.New("indigo: invalid character = " + string(id[i]))
		}
		n = n*fiftyEight + uint64(u)
	}
	return n, nil
}
