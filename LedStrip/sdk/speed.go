package sdk

import (
	"errors"
	"fmt"
)

var (
	ErrCouldNotSetSpeed = errors.New("Could not set speed")
)

type Speed struct {
	ls *Ledstrip
}

func (s *Speed) Get() (uint64, error) {
	panic("Speed/get - Not implemented yet")
	return 0, nil
}

func (s *Speed) Set(speed uint64) error {
	if !s.ls.ExecuteCommandBoolean(fmt.Sprintf("s %v", speed)) {
		return ErrCouldNotSetSpeed
	}

	s.ls.State.Speed = speed

	return nil
}
