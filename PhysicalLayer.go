package main

import (
	"errors"
	"fmt"
	"go.bug.st/serial.v1"
	"log"
	"time"
)

type PhysicalLayer struct {
	mode *serial.Mode
	portName string
	port serial.Port
}

const (
	CONNECT_TIMEOUT = 2 * time.Second
	SEND_TIMEOUT = 1 * time.Second
)

func (p *PhysicalLayer) connect() error {
	port, err := serial.Open("COM1", p.mode)
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = port.SetDTR(true)
	if err != nil {
		log.Fatal(err)
		return err
	}

	c1 := make(chan error, 1)
	go func() {
		defer close(c1)
		for {
			status, err := p.port.GetModemStatusBits()
			if err != nil {
				c1 <- err
				return
			}
			if status.DSR {
				c1 <- nil
				return
			}
		}
	}()

	select {
	case res := <-c1:
		if res == nil {
			p.port = port
		} else {
			return res
		}
	case <-time.After(CONNECT_TIMEOUT):
		fmt.Println("connect timeout 2")
		return errors.New("timeout error: DSR didnt get in 2 second")
	}

	p.port = port
	return nil
}

func (p *PhysicalLayer) disconnect() error {
	err := p.port.SetDTR(false)
	if err != nil {
		return err
	}
	err = p.port.Close()
	return err
}

func (p *PhysicalLayer) send(data []byte) (int, error) {
	err := p.port.SetRTS(true)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	c1 := make(chan error, 1)
	go func() {
		for {
			status, err := p.port.GetModemStatusBits()
			if err != nil {
				c1 <- err
				return
			}
			if status.CTS {
				c1 <- nil
				return
			}
		}
	}()

	select {
	case res := <-c1:
		if res == nil {
			n, err := p.port.Write(data)
			if err != nil {
				log.Fatal(err)
				return 0, err
			}
			return n, nil
		} else {
			return 0, res
		}
	case <-time.After(SEND_TIMEOUT):
		fmt.Println("send timeout 1")
		return 0, errors.New("timeout error: CTS didnt get in 1 second")
	}
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
