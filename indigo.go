package indigo

import (
	"errors"
	"sync"
	"time"

	"github.com/sony/sonyflake"
)

var (
	o sync.Once
	s *sonyflake.Sonyflake
)

// New settings new a sonyflake.
func New(start time.Time, machineID func() (uint16, error), checkMachineID func(uint16) bool) {
	o.Do(func() {
		s = sonyflake.NewSonyflake(sonyflake.Settings{
			StartTime:      start,
			MachineID:      machineID,
			CheckMachineID: checkMachineID,
		})
	})
}

// NextID generates a next unique ID.
func NextID() (string, error) {
	if s == nil {
		return "", errors.New("indigo: must be initialize a sonyflake")
	}
	n, err := s.NextID()
	if err != nil {
		return "", err
	}
	return EncodeBase58(n), nil
}

// Decompose returns a set of Sonyflake ID parts.
func Decompose(id string) (map[string]uint64, error) {
	b, err := DecodeBase58(id)
	if err != nil {
		return nil, err
	}
	return sonyflake.Decompose(b), nil
}
