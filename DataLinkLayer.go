package main

import (
	"go.bug.st/serial.v1"
	"log"
)

type DataLinkLayer struct {
	phys PhysicalLayer
}

func (d *DataLinkLayer) logicalConnect(mode *serial.Mode, portName string) (error) {
	d.phys = PhysicalLayer{mode: mode, portName: portName, port: nil }
	err := d.phys.connect()
	return err
}

func (d *DataLinkLayer) logicalDisconnect() error {
	err := d.phys.disconnect()
	return err
}

func (d *DataLinkLayer) sendInfoFrame(data []byte, to byte, from byte) error {
	frame, err := createFrame(INFO_FRAME, to, from, data)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return d.sendFrame(*frame)
}

func (d *DataLinkLayer) sendLinkFrameByMaster() error {
	// тут должен быть функционал для мастер компа (инициатора)
}

func (d *DataLinkLayer) sendLinkFrameBySlave() error {
	// тут должен быть функционал для слейв компа (инициатора)
}

func (d *DataLinkLayer) sendAckFrame(to byte, from byte) error {
	frame, err := createFrame(ACK_FRAME, to, from, nil)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return d.sendFrame(*frame)
}

func (d *DataLinkLayer) sendRetFrame(to byte, from byte) error {
	frame, err := createFrame(RET_FRAME, to, from, nil)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return d.sendFrame(*frame)
}

func (d *DataLinkLayer) sendUplinkFrame(to byte, from byte) error {
	frame, err := createFrame(UPLINK_FRAME, to, from, nil)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return d.sendFrame(*frame)
}

func (d *DataLinkLayer) sendFrame(frame Frame) error {
	_, err := d.phys.send(frame.toBytesArray())
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}


