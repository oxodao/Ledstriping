package services

import (
	"bufio"
	"encoding/json"
	"strings"
	"time"

	"github.com/oxodao/ledstrip/config"
	"github.com/oxodao/ledstrip/models"

	"github.com/tarm/serial"
)

type Provider struct {
	Config *config.Config

	SerialConfig *serial.Config
	Port         *serial.Port
	scanner      *bufio.Scanner

	CurrentState *models.State
}

/**
	@TODO Implement a way to map response to its original command
    Ex: CMD ARG MSG_ID
	    Resp: MSG_ID OK

	Wait for a certain timeout, if still no answer for the MSG_ID => Consider failed

	This will fix 500s when going too fast in the UI
**/

func (p *Provider) ReadLine() string {
	p.scanner.Scan()
	return strings.Trim(p.scanner.Text(), " \r\n\t")
}

func (p *Provider) ExecuteCommand(command string) (string, error) {
	command = strings.Trim(command, " \t\n\r")
	_, err := p.Port.Write([]byte(command + "\n"))
	if err != nil {
		return "", err
	}

	return p.ReadLine(), nil
}

func (p *Provider) ExecuteCommandBoolean(command string) bool {
	res, err := p.ExecuteCommand(command)
	if err != nil {
		return false
	}

	if res != "OK" {
		return false
	}

	return true
}

func (p *Provider) loadInitialState() (*models.State, error) {
	stt, err := p.ExecuteCommand("state")
	if err != nil {
		return nil, err
	}

	state := models.State{}
	err = json.Unmarshal([]byte(stt), &state)

	state.CleanColor()

	return &state, err
}

func NewProvider(cfg *config.Config) (*Provider, error) {
	c := &serial.Config{
		Name:        cfg.SerialPort,
		Baud:        9600,
		ReadTimeout: 1 * time.Second,
	}

	s, err := serial.OpenPort(c)
	if err != nil {
		panic(err)
	}

	prv := &Provider{
		Config:       cfg,
		SerialConfig: c,
		Port:         s,
		scanner:      bufio.NewScanner(s),
	}

	state, err := prv.loadInitialState()
	if err != nil {
		return nil, err
	}

	prv.CurrentState = state

	return prv, nil
}
