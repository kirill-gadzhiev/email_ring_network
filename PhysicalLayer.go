package main

import (
	"fmt"
	"go.bug.st/serial.v1"
	"log"
)

type PhysicalLayer struct {
	mode *serial.Mode
	portName string
	port serial.Port
}

func (p *PhysicalLayer) connect() error {
	port, err := serial.Open("COM1", p.mode)
	if err != nil {
		log.Fatal(err)
		return err
	}
	p.port = port
	return nil
}

func (p *PhysicalLayer) disconnect() error {
	err := p.port.Close()
	return err
}

func (p *PhysicalLayer) send(data []byte) (int, error) {
	n, err := p.port.Write(data)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	return n, nil
}

func (p *PhysicalLayer) listen() ([]byte, error) {
	buff := make([]byte, 100)
	var data []byte
	for {
		// Reads up to 100 bytes
		n, err := p.port.Read(buff)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		if n == 0 {
			fmt.Println("\nEOF")
			break
		} else {
			data = append(data, buff[:n]...)
			fmt.Printf("%v", string(buff[:n]))
		}
	}
	return data, nil
}

//func (p *PhysicalLayer) isAlive() bool {
//	p.port.
//	return true
//}

func printPorts() {
	// Retrieve the port list
	ports, err := serial.GetPortsList()
	if err != nil {
		fmt.Println("No serial ports found!")
		log.Fatal(err)
	}
	if len(ports) == 0 {
		fmt.Println("No serial ports found!")
		log.Fatal("No serial ports found!")
	}

	// Print the list of detected ports
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}
}
