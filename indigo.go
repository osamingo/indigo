package indigo

import (
	"time"

	"github.com/sony/sonyflake"
)

type (
	// A Generator has sonyflake, characters of Base58 and decode map.
	Generator struct {
		sonyflake  *sonyflake.Sonyflake
		base       []byte
		baseLength uint64
		decodes    []int64
	}

	// Settings has setting parameters for indigo.Generator.
	Settings struct {
		StartTime      time.Time
		MachineID      func() (uint16, error)
		CheckMachineID func(uint16) bool
		Base           []byte
	}
)

// New settings new a indigo.Generator.
func New(s Settings) *Generator {

	if len(s.Base) == 0 {
		s.Base = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")
	}

	return &Generator{
		sonyflake: sonyflake.NewSonyflake(sonyflake.Settings{
			StartTime:      s.StartTime,
			MachineID:      s.MachineID,
			CheckMachineID: s.CheckMachineID,
		}),
		base:       s.Base,
		baseLength: uint64(len(s.Base)),
		decodes:    defineDecodeMap(s.Base),
	}
}

// NextID generates a next unique ID.
func (g *Generator) NextID() (string, error) {
	n, err := g.sonyflake.NextID()
	if err != nil {
		return "", err
	}
	return g.EncodeBase58(n), nil
}

// Decompose returns a set of sonyflake ID parts.
func (g *Generator) Decompose(id string) (map[string]uint64, error) {
	b, err := g.DecodeBase58(id)
	if err != nil {
		return nil, err
	}
	return sonyflake.Decompose(b), nil
}
