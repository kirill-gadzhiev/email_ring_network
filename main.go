package main

import (
	"fmt"
	"go.bug.st/serial.v1"
	"log"
)

func main() {
	printPorts()

	mode := &serial.Mode{
		BaudRate: 9600,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}

	//testFrame, _ := createFrame(INFO_FRAME, BROADCAST, MAX_ADDRESS, []byte("EMAIL!RING!NETWORK!"))
	//bytesToSend := testFrame.toBytesArray()
	//fmt.Println("SENT: ", bytesToSend)
	//testSender(bytesToSend)

	slave(mode, "COM1", mode, "COM2")
	//master(mode, "COM2", mode, "COM1")
}


