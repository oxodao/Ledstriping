package sdk

import (
	"errors"
	"fmt"
)

var (
	ErrCouldNotSetMode = errors.New("Could not set mode")
	ErrUnknownMode     = errors.New("Unknown mode")
)

type Mode struct {
	ls *Ledstrip
}

func (m *Mode) GetAvailableModes() []string {
	return availableModes
}

func (m *Mode) Get() {
	panic("Mode/get - Not implemented yet")
}

func (m *Mode) Set(mode string) error {
	id := -1

	for i, j := range availableModes {
		if j == mode {
			id = i
			break
		}
	}

	if id == -1 {
		return ErrUnknownMode
	}

	if !m.ls.ExecuteCommandBoolean(fmt.Sprintf("m %v", id)) {
		return ErrCouldNotSetMode
	}

	m.ls.State.stateInternal.Mode = uint64(id)
	m.ls.State.Mode = mode

	return nil
}

var availableModes = []string{
	"Static",
	"Blink",
	"Breath",
	"Color Wipe",
	"Color Wipe Inverse",
	"Color Wipe Reverse",
	"Color Wipe Reverse Inverse",
	"Color Wipe Random",
	"Random Color",
	"Single Dynamic",
	"Multi Dynamic",
	"Rainbow",
	"Rainbow Cycle",
	"Scan",
	"Dual Scan",
	"Fade",
	"Theater Chase",
	"Theater Chase Rainbow",
	"Running Lights",
	"Twinkle",
	"Twinkle Random",
	"Twinkle Fade",
	"Twinkle Fade Random",
	"Sparkle",
	"Flash Sparkle",
	"Hyper Sparkle",
	"Strobe",
	"Strobe Rainbow",
	"Multi Strobe",
	"Blink Rainbow",
	"Chase White",
	"Chase Color",
	"Chase Random",
	"Chase Rainbow",
	"Chase Flash",
	"Chase Flash Random",
	"Chase Rainbow White",
	"Chase Blackout",
	"Chase Blackout Rainbow",
	"Color Sweep Random",
	"Running Color",
	"Running Red Blue",
	"Running Random",
	"Larson Scanner",
	"Comet",
	"Fireworks",
	"Fireworks Random",
	"Merry Christmas",
	"Fire Flicker",
	"Fire Flicker (soft)",
	"Fire Flicker (intense)",
	"Circus Combustus",
	"Halloween",
	"Bicolor Chase",
	"Tricolor Chase",
	"TwinkleFOX",
	"Custom 0",
	"Custom 1",
	"Custom 2",
	"Custom 3",
	"Custom 4",
	"Custom 5",
	"Custom 6",
	"Custom 7",
}
