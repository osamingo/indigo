package indigo

// An Encoder has Encode and Decode methods.
type Encoder interface {
	Encode(uint64) string
	Decode(string) (uint64, error)
}
