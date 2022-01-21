package sdk

import (
	"errors"
)

var (
	ErrCouldNotSetColor    = errors.New("Could not set color")
	ErrCouldNotSetDuration = errors.New("Could not set duration")
)

type Color struct {
	ls *Ledstrip
}

func (c *Color) Get() (string, error) {
	panic("Color/get - Not implemented yet")
	return "", nil
}

func (c *Color) Set(color string) error {
	if !c.ls.ExecuteCommandBoolean("c " + getBoardFromHex(color)) {
		return ErrCouldNotSetColor
	}

	// @TODO update internalstate
	c.ls.State.Color = color

	return nil
}

func (c *Color) Spark(color string, duration string) error {
	if !c.ls.ExecuteCommandBoolean("d " + duration) {
		return ErrCouldNotSetDuration
	}

	if !c.ls.ExecuteCommandBoolean("sp " + getBoardFromHex(color)) {
		return ErrCouldNotSetColor
	}

	return nil
}
