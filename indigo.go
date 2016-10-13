package indigo

import (
	"time"

	"github.com/osamingo/indigo/base58"
	"github.com/sony/sonyflake"
)

type (
	// A Generator has sonyflake, characters of BaseXX and decode map.
	Generator struct {
		sf  *sonyflake.Sonyflake
		enc Encoder
	}

	// Settings has setting parameters for indigo.Generator.
	Settings struct {
		StartTime      time.Time
		MachineID      func() (uint16, error)
		CheckMachineID func(uint16) bool
		Encoder        Encoder
	}
)

// New settings new a indigo.Generator.
func New(s Settings) *Generator {
	if s.Encoder == nil {
		s.Encoder = base58.StdEncoding
	}
	return &Generator{
		sf: sonyflake.NewSonyflake(sonyflake.Settings{
			StartTime:      s.StartTime,
			MachineID:      s.MachineID,
			CheckMachineID: s.CheckMachineID,
		}),
		enc: s.Encoder,
	}
}

// NextID generates a next unique ID.
func (g *Generator) NextID() (string, error) {
	n, err := g.sf.NextID()
	if err != nil {
		return "", err
	}
	return g.enc.Encode(n), nil
}

// Decompose returns a set of sonyflake ID parts.
func (g *Generator) Decompose(id string) (map[string]uint64, error) {
	b, err := g.enc.Decode(id)
	if err != nil {
		return nil, err
	}
	return sonyflake.Decompose(b), nil
}
