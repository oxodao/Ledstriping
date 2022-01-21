package sdk

import (
	"errors"
	"fmt"
)

var (
	ErrCouldNotSetBrightness = errors.New("Could not set brightness")
)

type Brightness struct {
	ls *Ledstrip
}

func (b *Brightness) Get() (uint64, error) {
	panic("Brightness/get - Not implemented yet")
	return 0, nil
}

func (b *Brightness) Set(brightness uint64) error {
	if !b.ls.ExecuteCommandBoolean(fmt.Sprintf("b %v", brightness)) {
		return ErrCouldNotSetBrightness
	}

	b.ls.State.Brightness = brightness
	b.ls.State.stateInternal.Brightness = brightness
	return nil
}

func (b *Brightness) FadeIn() error {
	if !b.ls.ExecuteCommandBoolean("fadein") {
		return ErrCouldNotSetBrightness
	}

	return nil
}

func (b *Brightness) FadeOut() error {
	if !b.ls.ExecuteCommandBoolean("fadeout") {
		return ErrCouldNotSetBrightness
	}

	return nil
}
