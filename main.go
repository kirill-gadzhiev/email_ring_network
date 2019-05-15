package main

import (
	"fmt"
	"go.bug.st/serial.v1"
	"log"
)

//func testSend() {
//	mode := &serial.Mode{
//		BaudRate: 9600,
//		Parity:   serial.NoParity,
//		DataBits: 8,
//		StopBits: serial.OneStopBit,
//	}
//
//	phys := &PhysicalLayer{mode: mode, portName: "COM1", port: nil }
//
//	err := phys.connect()
//	if err != nil {
//		fmt.Println("testSend ERROR" + err.Error())
//		return
//	}
//	defer phys.disconnect()
//
//	testData := []byte("Now we send some data")
//	_, err = phys.send(testData)
//
//	if err != nil {
//		fmt.Println("testSend ERROR" + err.Error())
//		return
//	}
//}
//
//func testRead() {
//	mode := &serial.Mode{
//		BaudRate: 9600,
//		Parity:   serial.NoParity,
//		DataBits: 8,
//		StopBits: serial.OneStopBit,
//	}
//
//	phys := &PhysicalLayer{mode: mode, portName: "COM1", port: nil }
//
//	err := phys.connect()
//	if err != nil {
//		fmt.Println("testRead ERROR" + err.Error())
//		return
//	}
//	defer phys.disconnect()
//
//	data, err := phys.listen()
//	if err != nil {
//		fmt.Println("testRead ERROR" + err.Error())
//		return
//	}
//
//	fmt.Println("Received data:" + string(data))
//}

func testSender() {
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

	testData := []byte("Now we send some data")
	_, _ = port.Write(testData)


}

func testReciever() {
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
	fmt.Println("RECEIVER OPENED PORT")
	for {
		status, _ := port.GetModemStatusBits()
		if status.DSR {
			fmt.Println("DSR:", status.DSR, " CTS:", status.CTS)

			err = port.SetDTR(true)
			if err != nil {
				log.Fatal(err)
				return
			}
			fmt.Println("DSR:", status.DSR, " CTS:", status.CTS)
			return
		}
	}

}



func main() {
	fmt.Println("Program starts")
	printPorts()
	//testRead()
	testSender()
	//testReciever()
}
