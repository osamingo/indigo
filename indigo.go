package indigo

import (
	"errors"
	"sort"
	"strings"
	"time"

	"github.com/sony/sonyflake"
)

type (
	// A Generator has sonyflake, characters of Base58 and decode map.
	Generator struct {
		sonyflake  *sonyflake.Sonyflake
		characters []byte
		decodeMap  []int64
	}

	// Settings has setting parameters for indigo.Generator.
	Settings struct {
		StartTime      time.Time
		MachineID      func() (uint16, error)
		CheckMachineID func(uint16) bool
		Characters     string
		Sort           bool
	}
)

// New settings new a indigo.Generator.
func New(s Settings) (*Generator, error) {

	if s.Characters == "" {
		s.Characters = defaultCharacters
	}

	if len(s.Characters) != fiftyEight {
		return nil, errors.New("indigo: characters must be 58 length")
	}

	if s.Sort {
		cs := strings.Split(s.Characters, "")
		sort.Strings(cs)
		s.Characters = strings.Join(cs, "")
	}

	g := &Generator{
		sonyflake: sonyflake.NewSonyflake(sonyflake.Settings{
			StartTime:      s.StartTime,
			MachineID:      s.MachineID,
			CheckMachineID: s.CheckMachineID,
		}),
		characters: []byte(s.Characters),
	}

	g.decodeMap = defineDecodeMap(g.characters)

	return g, nil
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
