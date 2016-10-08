package indigo

import (
	"encoding/binary"
	"errors"
)

func defineDecodeMap(characters []byte) []int64 {

	m := make([]int64, 256)

	for i := range m {
		m[i] = -1
	}
	for i := range characters {
		m[characters[i]] = int64(i)
	}

	return m
}

// EncodeBase58 returns encoded string by Base58.
func (g *Generator) EncodeBase58(id uint64) string {

	if id == 0 {
		return string(g.base[:1])
	}

	bin := make([]byte, 0, binary.MaxVarintLen64)
	for id > 0 {
		bin = append(bin, g.base[id%g.baseLength])
		id /= g.baseLength
	}

	for i, j := 0, len(bin)-1; i < j; i, j = i+1, j-1 {
		bin[i], bin[j] = bin[j], bin[i]
	}

	return string(bin)
}

// DecodeBase58 returns decoded unsigned int64 by Base58.
func (g *Generator) DecodeBase58(id string) (uint64, error) {

	if id == "" {
		return 0, errors.New("indigo: source should not be empty")
	}

	var n uint64
	for i := range id {
		u := g.decodes[id[i]]
		if u < 0 {
			return 0, errors.New("indigo: invalid character = " + string(id[i]))
		}
		n = n*g.baseLength + uint64(u)
	}
	return n, nil
}
