package main

import (
	"encoding/json"
	"fmt"
	"go.bug.st/serial.v1"
	"log"
	"strings"
)

type DataLinkLayer struct {
	phys PhysicalLayer
	state map[byte]string
	currentAddress byte
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

func (d *DataLinkLayer) sendPingFrame(from byte) error {
	frame, err := createFrame(INFO_FRAME, BROADCAST, from, nil)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return d.sendFrame(*frame)
}

func (d *DataLinkLayer) sendLinkFrameByMaster(email string) error {
	// тут должен быть функционал для мастер компа (инициатора)

	fmt.Println("1")
	data := []byte(string(MIN_ADDRESS) + " " + email + "\n")

	fmt.Println("2")
	frame, err := createFrame(LINK_FRAME, BROADCAST, MIN_ADDRESS, data)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
		return err
	}
	fmt.Println("3")

	d.currentAddress = MIN_ADDRESS

	fmt.Println("4")
	return d.sendFrame(*frame)
}

// возвращает адрес текущего компа
func (d *DataLinkLayer) sendLinkFrameBySlave(receivedFrame Frame, email string) error {
	// тут должен быть функционал для слейв компа

	//strData := string(receivedFrame.data)
	//clients := strings.Split(strData, "\n")
	//last := clients[len(clients) - 1]
	//lastData := strings.Split(last, " ")
	//lastAddress := lastData[0][0]  // адрес однобайтовый, значит больше одного символа в строке занимать не может
	currentAddress := receivedFrame.from + 1
	currentData := []byte(string(currentAddress) + " " + email + "\n")

	newData := append(receivedFrame.data, currentData...)

	frame, err := createFrame(LINK_FRAME, BROADCAST, currentAddress, newData)
	if err != nil {
		log.Fatal(err)
		return err
	}

	d.currentAddress = currentAddress

	return d.sendFrame(*frame)
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

func (d *DataLinkLayer) terminate() error {
	fmt.Println("TERMINATE")
	err := d.sendUplinkFrame(BROADCAST, d.currentAddress)
	if err != nil {
		return err
	}
	return nil
}

// эту функцию будем пихать в горутину наверн и она будет писать в канал если что-то пришло
// в основном потоке считываем и действуем в зависимости от этого
func (d *DataLinkLayer) listen(out chan<- []byte) error {
	for {
		fmt.Println("inside listen for")
		// Reads up to 100 bytes
		buff := make([]byte, 100)
		n, err := d.phys.port.Read(buff)
		if err != nil {
			continue
			return err
		}
		out <- buff[:n]
	}
}

func (d *DataLinkLayer) sendStateFrame(state []byte, from byte) error {
	frame, err := createFrame(STATE_FRAME, BROADCAST, from, state)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return d.sendFrame(*frame)
}

func (d *DataLinkLayer) sendLetter(letter Letter) {
	var to byte
	for address, email := range d.state {
		if email == letter.Responder {
			to = address
		}
	}
	data, err := json.Marshal(letter)
	if err != nil {
		fmt.Println("letter marshalling error")
		return
	}
	d.sendInfoFrame(data, to, d.currentAddress)
}

func (d *DataLinkLayer) getStateFromFrame(receivedFrame Frame) (map[byte]string) {
	strData := string(receivedFrame.data)
	dataLen := len(strData)
	trimmedData := strData[:dataLen-1]
	clientsData := strings.Split(trimmedData, "\n")

	clients := map[byte]string{}

	for _, client := range clientsData {
		clientInfo := strings.Split(client, " ")

		// адрес однобайтовый, значит больше одного символа в строке занимать не может
		fmt.Println("client: ", client, " clientInfo: ", clientInfo)
		clientAddress := byte(clientInfo[0][0])
		clientEmail := clientInfo[1]

		clients[clientAddress] = clientEmail
	}
	return clients
}





