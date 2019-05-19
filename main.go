package main

import (
	"fmt"
	"go.bug.st/serial.v1"
	"log"
)

func testSender(data []byte) {
	mode := &serial.Mode{
		BaudRate: 9600,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}

	port, err := serial.Open("COM1", mode)
	if err != nil {
		log.Fatal(err)
		return
	}


	//var a string
	//_, _ = fmt.Scan(a)
	fmt.Println("SENDER OPENED PORT")
	status, _ := port.GetModemStatusBits()
	fmt.Println("DSR:", status.DSR, " CTS:", status.CTS)

	err = port.SetDTR(false)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("DSR:", status.DSR, " CTS:", status.CTS)

	_, _ = port.Write(data)


}

func main() {
	init_interface()
	fmt.Println("Program starts MASTER")
	//fmt.Println("Program starts SLAVE")
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

	//slave(mode, "COM1", mode, "COM2")
	master(mode, "COM2", mode, "COM1")
}


