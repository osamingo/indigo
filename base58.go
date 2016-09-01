package indigo

import (
	"errors"
	"math/big"
)

const alphanumeric = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

var decodeMap [256]int64

func init() {
	for i := range decodeMap {
		decodeMap[i] = -1
	}
	for i := range alphanumeric {
		decodeMap[alphanumeric[i]] = int64(i)
	}
}

// EncodeBase58 returns encoded string by Base58.
func EncodeBase58(u uint64) string {

	if u == 0 {
		return "1"
	}

	n := new(big.Int).SetUint64(u)

	mod := new(big.Int)
	zero := big.NewInt(0)
	radix := big.NewInt(58)
	dst := make([]byte, 0, len(n.Bytes()))
	for n.Cmp(zero) > 0 {
		n.DivMod(n, radix, mod)
		dst = append(dst, alphanumeric[mod.Int64()])
	}

	for i, j := n.BitLen(), len(dst)-1; i < j; i, j = i+1, j-1 {
		dst[i], dst[j] = dst[j], dst[i]
	}

	return string(dst)
}

// DecodeBase58 returns decoded unsigned int64 by Base58.
func DecodeBase58(s string) (uint64, error) {

	if s == "" {
		return 0, errors.New("indigo: source should not be empty")
	}

	n := new(big.Int)
	radix := big.NewInt(58)
	u := int64(0)
	for i := range s {
		if u = decodeMap[s[i]]; u < 0 {
			return 0, errors.New("indigo: invalid character = " + string(s[i]))
		}
		n.Add(n.Mul(n, radix), big.NewInt(u))
	}
	return n.Uint64(), nil
}
