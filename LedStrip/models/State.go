package models

import (
	"encoding/json"
	"fmt"
)

type State struct {
	DirtyColor         uint64 `json:"color"`
	Mode               uint64 `json:"mode"`
	Brightness         uint64 `json:"brightness"`
	Speed              uint64 `json:"speed"`
	Color              string `json:"color_clean"`
	Sparking           bool   `json:"sparking"`
	Spark_start        uint64 `json:"spark_start"`
	Spark_duration     uint64 `json:"spark_duration"`
	Fading             uint64 `json:"fading"`
	OriginalBrightness uint64 `json:"original_brightness"`
}

func (s *State) CleanColor() {
	s.Color = fmt.Sprintf("0x%02X%04X", s.DirtyColor>>16, s.DirtyColor&0xFFFF)
}

func (s State) Json() []byte {
	//s.CleanColor()

	state := struct {
		Mode       uint64
		Brightness uint64
		Speed      uint64
		Color      string
	}{
		Mode:       s.Mode,
		Brightness: s.Brightness,
		Speed:      s.Speed,
		Color:      s.Color,
	}

	str, _ := json.Marshal(state)

	return str
}
