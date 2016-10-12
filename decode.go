package indigo

import "errors"

// Decode returns decoded unsigned int64 by BaseXX.
func (g *Generator) Decode(id string) (uint64, error) {

	if id == "" {
		return 0, errors.New("indigo: source should not be empty")
	}

	n, l := uint64(0), uint64(len(g.base))
	for i := range id {
		u := g.decodes[id[i]]
		if u < 0 {
			return 0, errors.New("indigo: invalid character = " + string(id[i]))
		}
		n = n*l + uint64(u)
	}
	return n, nil
}

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
