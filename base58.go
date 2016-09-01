package indigo

import "errors"

const alphanumeric = "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"

var decodeMap = make([]int64, 256)

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

	d := make([]byte, 0, 10)
	for u > 0 {
		d = append(d, alphanumeric[u%58])
		u = u / 58
	}

	for i, j := 0, len(d)-1; i < j; i, j = i+1, j-1 {
		d[i], d[j] = d[j], d[i]
	}

	return string(d)
}

// DecodeBase58 returns decoded unsigned int64 by Base58.
func DecodeBase58(s string) (uint64, error) {

	if s == "" {
		return 0, errors.New("indigo: source should not be empty")
	}

	n := uint64(0)
	for i := range s {
		u := decodeMap[s[i]]
		if u < 0 {
			return 0, errors.New("indigo: invalid character = " + string(s[i]))
		}
		n = n*58 + uint64(u)
	}
	return n, nil
}
