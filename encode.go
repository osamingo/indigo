package indigo

import "encoding/binary"

// Encode returns encoded string by BaseXX.
func (g *Generator) Encode(id uint64) string {

	if id == 0 {
		return string(g.base[:1])
	}

	l, bin := uint64(len(g.base)), make([]byte, 0, binary.MaxVarintLen64)
	for id > 0 {
		bin = append(bin, g.base[id%l])
		id /= l
	}

	for i, j := 0, len(bin)-1; i < j; i, j = i+1, j-1 {
		bin[i], bin[j] = bin[j], bin[i]
	}

	return string(bin)
}
