package obd

import (
	"obd2-tool/modes"
	"obd2-tool/utils"
	"strings"
)

func (c *Client) ReadEngineSpeed() (float64, error) {
	raw, err := c.PidExec(modes.ModeCurrent, modes.PID(modes.PidEngineSpeed))
	if err != nil {
		return 0, err
	}

	parts := strings.Split(raw, " ")
	data, err := utils.HexToIntArray(parts[len(parts)-2:])
	if err != nil {
		return 0, err
	}

	return float64(256.*data[0]+data[1]) / 4., nil
}

func (c *Client) ReadThrottlePosition() (float64, error) {
	raw, err := c.PidExec(modes.ModeCurrent, modes.PID(modes.PidThrottlePosition))
	if err != nil {
		return 0, err
	}

	parts := strings.Split(raw, " ")
	data, err := utils.HexToIntArray(parts[len(parts)-1:])
	if err != nil {
		return 0, err
	}

	return 100. / 255. * float64(data[0]), nil
}

func (c *Client) ReadEngineCoolantTemp() (int, error) {
	raw, err := c.PidExec(modes.ModeCurrent, modes.PID(modes.PidEngineCoolantTemp))
	if err != nil {
		return 0, err
	}

	parts := strings.Split(raw, " ")
	data, err := utils.HexToIntArray(parts[len(parts)-1:])
	if err != nil {
		return 0, err
	}

	return data[0] - 40, nil
}

func (c *Client) ReadMassAirflow() (float64, error) {
	raw, err := c.PidExec(modes.ModeCurrent, modes.PID(modes.PidMassAirflow))
	if err != nil {
		return 0, err
	}

	parts := strings.Split(raw, " ")
	data, err := utils.HexToIntArray(parts[len(parts)-2:])
	if err != nil {
		return 0, err
	}

	return float64(256.*data[0]+data[1]) / 100., nil
}

func (c *Client) ReadVehicleSpeed() (int, error) {
	raw, err := c.PidExec(modes.ModeCurrent, modes.PID(modes.PidVehicleSpeed))
	if err != nil {
		return 0, err
	}

	parts := strings.Split(raw, " ")
	data, err := utils.HexToIntArray(parts[len(parts)-2:])
	if err != nil {
		return 0, err
	}

	return data[0], nil
}
