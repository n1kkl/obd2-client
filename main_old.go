package main

import (
	"bufio"
	"github.com/tarm/serial"
	"log"
	"math/big"
	"strings"
	"time"
)

func main_old() {
	port, err := serial.OpenPort(&serial.Config{
		Name:        "CNCB0",
		Baud:        38400,
		ReadTimeout: time.Millisecond * 500,
	})
	if err != nil {
		log.Fatalf("failed to connect to serial port: %v", err)
	}

	commands := []string{"ATZ", "ATD", "ATSP6", "ATE0"}
	reader := bufio.NewReader(port)

	for _, cmd := range commands {
		log.Println("->", cmd)
		err := send(port, cmd)
		if err != nil {
			log.Fatalf("failed to write to serial port: %v", err)
		}

		res, err := read(reader)
		if err != nil {
			log.Fatalf("failed to read from serial port: %v", err)
		}

		log.Println("<-", res)
	}

	for true {
		err := send(port, "010D")
		if err != nil {
			log.Fatalf("failed to write to serial port: %v", err)
		}

		res, err := read(reader)
		if err != nil {
			log.Fatalf("failed to read from serial port: %v", err)
		}

		bytes := strings.Split(res, " ")
		speed := new(big.Int)
		speed.SetString(bytes[2], 16)
		log.Println(speed)
	}
}

func send(port *serial.Port, cmd string) error {
	_, err := port.Write([]byte(cmd + "\r"))
	if err != nil {
		return err
	}
	return nil
}

func read(reader *bufio.Reader) (string, error) {
	data, err := reader.ReadBytes('>')
	if err != nil {
		return "", err
	}
	return strings.Trim(string(data[:]), "\r\n>"), nil
}
