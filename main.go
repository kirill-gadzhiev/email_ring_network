package main

import (
	"fmt"
	"go.bug.st/serial.v1"
)

func testSend() {
	mode := &serial.Mode{
		BaudRate: 9600,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}

	phys := &PhysicalLayer{mode: mode, portName: "COM1", port: nil }

	err := phys.connect()
	if err != nil {
		fmt.Println("testSend ERROR" + err.Error())
		return
	}
	defer phys.disconnect()

	testData := []byte("Now we send some data")
	_, err = phys.send(testData)

	if err != nil {
		fmt.Println("testSend ERROR" + err.Error())
		return
	}
}

func testRead() {
	mode := &serial.Mode{
		BaudRate: 9600,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}

	phys := &PhysicalLayer{mode: mode, portName: "COM1", port: nil }

	err := phys.connect()
	if err != nil {
		fmt.Println("testRead ERROR" + err.Error())
		return
	}
	defer phys.disconnect()

	data, err := phys.listen()
	if err != nil {
		fmt.Println("testRead ERROR" + err.Error())
		return
	}

	fmt.Println("Received data:" + string(data))
}

func main()  {
	fmt.Println("Program starts")
	printPorts()
	testRead()
}
