package sdk

import (
	"bufio"
	"strings"
	"sync"
	"time"

	"github.com/tarm/serial"
)

type Ledstrip struct {
	serialConfig *serial.Config
	serialPort   *serial.Port
	scanner      *bufio.Scanner

	State *State

	Brightness *Brightness
	Color      *Color
	Mode       *Mode
	Speed      *Speed

	cmdMutex sync.Mutex
}

func Connect(port string) (*Ledstrip, error) {
	c := &serial.Config{
		Name:        port,
		Baud:        9600,
		ReadTimeout: 1 * time.Second,
	}

	s, err := serial.OpenPort(c)
	if err != nil {
		return nil, err
	}

	ls := &Ledstrip{
		serialConfig: c,
		serialPort:   s,
		scanner:      bufio.NewScanner(s),
	}

	ls.Brightness = &Brightness{ls}
	ls.Color = &Color{ls}
	ls.Mode = &Mode{ls}
	ls.Speed = &Speed{ls}
	ls.State = &State{ls: ls}

	err = ls.State.Fetch()
	if err != nil {
		return nil, err
	}

	return ls, nil
}

func (ls *Ledstrip) readLine() string {
	ls.scanner.Scan()
	return strings.Trim(ls.scanner.Text(), " \r\n\t")

}

func (ls *Ledstrip) ExecuteCommand(command string) (string, error) {
	ls.cmdMutex.Lock()

	command = strings.Trim(command, " \t\n\r")
	_, err := ls.serialPort.Write([]byte(command + "\n"))
	if err != nil {
		return "", err
	}

	resp := ls.readLine()

	ls.cmdMutex.Unlock()

	return resp, nil
}

func (ls *Ledstrip) ExecuteCommandBoolean(command string) bool {
	res, err := ls.ExecuteCommand(command)
	if err != nil {
		return false
	}

	if res != "OK" {
		return false
	}

	return true
}

func (ls *Ledstrip) SetState(color string, brightness uint64, mode string, speed uint64) error {
	if err := ls.Color.Set(color); err != nil {
		return err
	}

	if err := ls.Brightness.Set(brightness); err != nil {
		return err
	}

	if err := ls.Mode.Set(mode); err != nil {
		return err
	}

	if err := ls.Speed.Set(speed); err != nil {
		return err
	}

	return nil
}

// @TODO: Fix the firmware to process colors correctly ?

func getHexFromBoard(boardColor string) string {
	return "#" + boardColor[2:4] + boardColor[6:8] + boardColor[4:6]
}

func getBoardFromHex(hexColor string) string {
	if len(hexColor) < 7 {
		return "000000"
	}

	return "0x" + hexColor[1:3] + hexColor[5:7] + hexColor[3:5]
}
