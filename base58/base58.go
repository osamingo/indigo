package base58

import (
	"encoding/binary"
	"errors"
)

// An Encoder implements indigo.Encoder interface by Base58.
type Encoder struct {
	encode    [58]byte
	decodeMap [256]int
}

// StdEncoding is Base58 Encoder.
var StdEncoding = MustNewEncoder("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

// MustNewEncoder returns new base58.Encoder.
func MustNewEncoder(source string) *Encoder {
	enc, err := NewEncoder(source)
	if err != nil {
		panic(err)
	}
	return enc
}

// NewEncoder returns new base58.Encoder.
func NewEncoder(source string) (*Encoder, error) {

	if len(source) != 58 {
		return nil, errors.New("base58: encoding source is not 58-bytes long")
	}

	enc := new(Encoder)

	for i := range enc.decodeMap {
		enc.decodeMap[i] = -1
	}

	for i := range source {
		enc.encode[i] = source[i]
		enc.decodeMap[enc.encode[i]] = i
	}

	return enc, nil
}

// Encode returns encoded string by Base58.
func (enc *Encoder) Encode(id uint64) string {

	if id == 0 {
		return string(enc.encode[:1])
	}

	bin := make([]byte, 0, binary.MaxVarintLen64)
	for id > 0 {
		bin = append(bin, enc.encode[id%58])
		id /= 58
	}

	for i, j := 0, len(bin)-1; i < j; i, j = i+1, j-1 {
		bin[i], bin[j] = bin[j], bin[i]
	}

	return string(bin)
}

// Decode returns decoded unsigned int64 by Base58.
func (enc *Encoder) Decode(id string) (uint64, error) {

	if id == "" {
		return 0, errors.New("base58: id should not be empty")
	}

	var n uint64
	for i := range id {
		u := enc.decodeMap[id[i]]
		if u < 0 {
			return 0, errors.New("base58: invalid character - " + string(id[i]))
		}
		n = n*58 + uint64(u)
	}

	return n, nil
}
