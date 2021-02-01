/*
Package indigo generates a distributed unique ID generator of using Sonyflake and encoded by Base58.

More information: https://github.com/osamingo/indigo/blob/master/README.md
*/
package indigo

import (
	"fmt"
	"time"

	"github.com/osamingo/base58"
	"github.com/sony/sonyflake"
)

type (
	// An Encoder has Encode and Decode methods.
	Encoder interface {
		Encode(uint64) string
		Decode(string) (uint64, error)
	}
	// A Generator has sonyflake and encoder.
	Generator struct {
		sf  *sonyflake.Sonyflake
		enc Encoder
	}
)

// New settings new a indigo.Generator.
func New(enc Encoder, options ...func(*sonyflake.Settings)) *Generator {
	if enc == nil {
		enc = base58.MustNewEncoder(base58.StandardSource)
	}

	s := sonyflake.Settings{}

	for i := range options {
		options[i](&s)
	}

	return &Generator{
		sf:  sonyflake.NewSonyflake(s),
		enc: enc,
	}
}

// StartTime is optional function for indigo.Generator.
func StartTime(t time.Time) func(*sonyflake.Settings) {
	return func(s *sonyflake.Settings) {
		s.StartTime = t
	}
}

// MachineID is optional function for indigo.Generator.
func MachineID(f func() (uint16, error)) func(*sonyflake.Settings) {
	return func(s *sonyflake.Settings) {
		s.MachineID = f
	}
}

// CheckMachineID is optional function for indigo.Generator.
func CheckMachineID(f func(uint16) bool) func(*sonyflake.Settings) {
	return func(s *sonyflake.Settings) {
		s.CheckMachineID = f
	}
}

// NextID generates a next unique ID.
func (g *Generator) NextID() (string, error) {
	n, err := g.sf.NextID()
	if err != nil {
		return "", fmt.Errorf("indigo: failed to generate next id: %w", err)
	}

	return g.enc.Encode(n), nil
}

// Decompose returns a set of sonyflake ID parts.
func (g *Generator) Decompose(id string) (map[string]uint64, error) {
	b, err := g.enc.Decode(id)
	if err != nil {
		return nil, fmt.Errorf("indigo: failed to decode, id = %s: %w", id, err)
	}

	return sonyflake.Decompose(b), nil
}
