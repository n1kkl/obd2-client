package main

import (
	"fmt"
	"github.com/tarm/serial"
	"log"
	"obd2-tool/obd"
	"time"
)

func main() {
	client, err := obd.NewClient(&serial.Config{
		Name:        "COM4",
		Baud:        38400,
		ReadTimeout: time.Millisecond * 500,
	})
	if err != nil {
		log.Fatalf("could not initialize connection with serial device: %v", err)
	}
	//goland:noinspection GoUnhandledErrorResult
	defer client.Close()

	log.Println("successfully connected to serial device")

	for true {
		res, err := client.ReadVehicleSpeed()
		if err != nil {
			log.Fatalf("failed to execute pid command: %v", err)
		}
		fmt.Println(res)
		time.Sleep(time.Millisecond * 500)
	}
}
