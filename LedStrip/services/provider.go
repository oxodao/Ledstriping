package services

import (
	"bufio"
	"errors"
	"github.com/oxodao/ledstrip/config"
	"github.com/oxodao/ledstrip/models"
	"strings"
	"time"

	"github.com/tarm/serial"
)

type Provider struct {
	Config *config.Config

	SerialConfig *serial.Config
	Port         *serial.Port
	scanner      *bufio.Scanner

	CurrentState *models.State
}

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
	color, err := p.ExecuteCommand("c")
	brightness, err2 := p.ExecuteCommand("b")
	mode, err3 := p.ExecuteCommand("m")
	speed, err4 := p.ExecuteCommand("s")

	if err != nil || err2 != nil || err3 != nil || err4 != nil {
		return nil, errors.New("something went wrong reading initial state")
	}

	state, err := models.NewState(mode, brightness, speed, color)
	if err != nil {
		return nil, errors.New("something went wrong parsing the state")
	}

	return state, nil
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
		Config: cfg,
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
