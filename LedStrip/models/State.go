package models

import (
	"github.com/oxodao/ledstrip/utils"
	"strconv"
)

type State struct {
	Mode uint64
	Brightness uint64
	Speed uint64
	Color string
}

func NewState(mode, brightness, speed, color string) (*State, error) {
	modeInt, err := strconv.ParseUint(mode, 10, 64)
	if err != nil {
		return nil, err
	}

	brightnessInt, err := strconv.ParseUint(brightness, 10, 64)
	if err != nil {
		return nil, err
	}

	speedInt, err := strconv.ParseUint(speed, 10, 64)
	if err != nil {
		return nil, err
	}

	return &State{
		Mode: modeInt,
		Brightness: brightnessInt,
		Speed: speedInt,
		Color: utils.GetHexFromBoard(color),
	}, nil
}