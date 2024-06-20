package obd

import (
	"bufio"
	"github.com/tarm/serial"
	"obd2-tool/commands"
)

// Client represents a serial connection to an OBD-II device.
type Client struct {
	port   *serial.Port
	reader *bufio.Reader
	remote *Remote
}

func NewClient(config *serial.Config) (*Client, error) {
	port, err := serial.OpenPort(config)
	if err != nil {
		return nil, err
	}

	client := &Client{
		port:   port,
		reader: bufio.NewReader(port),
	}
	err = client.Initialize()
	if err != nil {
		return client, err
	}
	return client, nil
}

func (c *Client) Initialize() error {
	initSequence := []commands.AtCommand{
		commands.AtWarmStart,
		commands.AtResetAll,
		commands.AtEchoOff,
		commands.AtLineFeedsOff,
		commands.AtHeadersOn,
	}

	for _, cmd := range initSequence {
		_, err := c.RawExec(string(cmd))
		if err != nil {
			return err
		}
	}

	devDescription, err := c.RawExec(commands.AtDeviceDescription)
	if err != nil {
		return err
	}

	devIdentifier, err := c.RawExec(commands.AtDeviceIdentifier)
	if err != nil {
		return err
	}

	c.remote = &Remote{
		devDescription,
		devIdentifier,
	}

	return nil
}
