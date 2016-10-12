package indigo

// A Encoder has Encode and Decode methods.
type Encoder interface {
	Encode(id uint64) string
	Decode(id string) (uint64, error)
}
