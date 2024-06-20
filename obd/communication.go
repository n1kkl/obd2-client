package obd

import (
	"fmt"
	"obd2-tool/commands"
	"obd2-tool/modes"
	"strings"
)

// PidExec executes the given pid command in a specific mode and returns the raw response
func (c *Client) PidExec(mode modes.Mode, pid modes.PID) (string, error) {
	return c.RawExec(fmt.Sprintf("%.2X%.2X", mode, pid))
}

// RawExec sends commands and immediately reads and returns the response
func (c *Client) RawExec(cmd string) (string, error) {
	_, err := c.Send(cmd)
	if err != nil {
		return "", err
	}

	return c.Read()
}

// Send writes raw command and appends carriage return
func (c *Client) Send(cmd string) (int, error) {
	return c.port.Write([]byte(cmd + "\r"))
}

// Read reads from the buffered reader until it sees the delimiter
func (c *Client) Read() (string, error) {
	data, err := c.reader.ReadBytes('>')
	if err != nil {
		return "", err
	}
	return strings.Trim(string(data[:]), "\n>"), nil
}

// Close sends reset all command (ATZ) and closes the connection
func (c *Client) Close() error {
	_, err := c.Send(commands.AtResetAll)
	if err != nil {
		return err
	}
	return c.port.Close()
}
