package main

import (
	"flag"
	"go.bug.st/serial.v1"
)

func main() {
	printPorts()

	in := flag.String("in", "COM1", "входящий COM-порт")
	out := flag.String("out", "COM2", "исходящий COM-порт")
	isMaster := flag.Bool("master", false, "запустить версию ПО для администратора")

	mode := &serial.Mode{
		BaudRate: 9600,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}

	flag.Parse()

	if *isMaster {
		master(mode, *in, mode, *out)
	} else {
		slave(mode, *in, mode, *out)
	}
}


