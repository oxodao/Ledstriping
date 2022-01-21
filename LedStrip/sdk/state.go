package sdk

import (
	"encoding/json"
	"fmt"
)

var (
	ErrCouldNotGetState   = fmt.Errorf("could not get state")
	ErrCouldNotParseState = fmt.Errorf("could not parse state")
)

type stateInternal struct {
	Mode               uint64 `json:"mode"`
	Brightness         uint64 `json:"brightness"`
	Speed              uint64 `json:"speed"`
	Color              uint64 `json:"color"`
	Sparking           bool   `json:"sparking"`
	Spark_start        uint64 `json:"spark_start"`
	Spark_duration     uint64 `json:"spark_duration"`
	Fading             uint64 `json:"fading"`
	OriginalBrightness uint64 `json:"original_brightness"`
}

type State struct {
	ls            *Ledstrip     `json:"-"`
	stateInternal stateInternal `json:"-"`

	Mode       string `json:"Mode"`
	Brightness uint64 `json:"Brightness"`
	Speed      uint64 `json:"Speed"`
	Color      string `json:"Color"`
}

func (s *State) updateFromInternal() {
	s.Mode = s.ls.Mode.GetAvailableModes()[s.stateInternal.Mode]
	s.Color = fmt.Sprintf("0x%02X%04X", s.stateInternal.Color>>16, s.stateInternal.Color&0xFFFF)
	s.Speed = s.stateInternal.Speed
	s.Brightness = s.stateInternal.Brightness
}

func (s *State) Fetch() error {
	state, err := s.ls.ExecuteCommand("state")
	if err != nil {
		return ErrCouldNotGetState
	}

	err = json.Unmarshal([]byte(state), &s.stateInternal)
	if err != nil {
		return ErrCouldNotParseState
	}

	s.updateFromInternal()

	return nil
}
