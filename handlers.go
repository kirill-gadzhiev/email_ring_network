package main

import (
	"errors"
	"fmt"
	"go.bug.st/serial.v1"
)

func master(backPortMode *serial.Mode, backPortName string, frontPortMode *serial.Mode, frontPortName string) error {
	// 0. Подключиться к заднему и переднему портам
	var backPort DataLinkLayer
	err := backPort.logicalConnect(backPortMode, backPortName)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var frontPort DataLinkLayer
	err = frontPort.logicalConnect(frontPortMode, frontPortName)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// 1. Отправить LINK_FRAME
	err = frontPort.sendLinkFrameByMaster()
	if err != nil {
		fmt.Println(err)
		return err
	}

	// 2. Дождаться получения LINK_FRAME после круга по кольцу
	ch := make(chan []byte, 3)
	go backPort.listen(ch)

	linkFrame, err := parseBytesToFrame(<-ch)
	if err != nil {
		fmt.Println(err)
	}
	if linkFrame.frameType != LINK_FRAME {
		// тут по хорошему надо делать что-то
		// ситуация, если кто-то прислал кадр то инициализации кольца
	}

	// думаю, что здесь будет передача стейта в горутину с интерфейсом, больше она почти нигде не нужна
	frontPort.state = frontPort.getStateFromFrame(*linkFrame)  // сохранить state куда-то

	// 3. Отправить STATE_FRAME
	err = frontPort.sendStateFrame(linkFrame.data, MIN_ADDRESS)
	if err != nil {
		fmt.Println(err)
	}


	// 4. Дождаться получения STATE_FRAME
	_, err = parseBytesToFrame(<-ch)
	if err != nil {
		fmt.Println(err)
	}

	// 5. Стать Slave
	// someLogicHere()
	for {
		frame, err := parseBytesToFrame(<-ch)
		if err != nil {
			fmt.Println(err)
		}

		err = frontPort.handleFrame(*frame)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func slave(backPortMode *serial.Mode, backPortName string, frontPortMode *serial.Mode, frontPortName string) error {
	// 0. Подключиться к заднему и переднему портам
	var backPort DataLinkLayer
	err := backPort.logicalConnect(backPortMode, backPortName)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var frontPort DataLinkLayer
	err = frontPort.logicalConnect(frontPortMode, frontPortName)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// 1. Слушаем задний
	ch := make(chan []byte, 3)
	go backPort.listen(ch)

	// 2. Парсим фрейм
	// 3. switch-case по frameType
	for {
		frame, err := parseBytesToFrame(<-ch)
		if err != nil {
			fmt.Println(err)
		}

		err = frontPort.handleFrame(*frame)
		if err != nil {
			fmt.Println(err)
		}
	}
}


func (d *DataLinkLayer) handleFrame(receivedFrame Frame) error {

	// здесь мы пересылаем кадр если он не предназначен нам
	if receivedFrame.to != d.currentAddress {
		if receivedFrame.to != BROADCAST {
			return d.sendFrame(receivedFrame)
		}
	}

	switch receivedFrame.frameType {
	case INFO_FRAME:
		return d.handleInfoFrame(receivedFrame)
	case LINK_FRAME:
		return d.handleLinkFrame(receivedFrame)
	case UPLINK_FRAME:
		return d.handleUplinkFrame(receivedFrame)
	case ACK_FRAME:
		return d.handleAckFrame(receivedFrame)
	case RET_FRAME:
		return d.handleRetFrame(receivedFrame)
	case STATE_FRAME:
		return d.handleStateFrame(receivedFrame)
	default:
		return errors.New("unknown type of Frame")
	}
	return nil
}

func (d *DataLinkLayer) handleInfoFrame(receivedFrame Frame) error {
	// 1. обработать фрейм

	// отправляем в горутину с интерфейсом инфу

	// 2. отправить ack
	err := d.sendAckFrame(receivedFrame.from, receivedFrame.to)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if receivedFrame.to == BROADCAST {
		err = d.sendFrame(receivedFrame)
		if err != nil {
			fmt.Println(err)
		}
	}

	return err
}

func (d *DataLinkLayer) handleLinkFrame(receivedFrame Frame) error {
	// 1. сохранить свой адрес (теперь он автоматически сохраняется при отправке кадра)
	// 2. отправить link далее
	email := ""  // тут надо сходить в интерфейс(?) и получить почту
	err := d.sendLinkFrameBySlave(receivedFrame, email)

	return err
}

func (d *DataLinkLayer) handleUplinkFrame(receivedFrame Frame) error {
	// 1. отправляем uplink далее
	err := d.sendFrame(receivedFrame)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// 2. разрываем соединение
	err = d.logicalDisconnect()
	return err
}

func (d *DataLinkLayer) handleAckFrame(receivedFrame Frame) error {
	// 1. ничего не делаем (?)
	if receivedFrame.to != d.currentAddress {
		err := d.sendFrame(receivedFrame)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}

func (d *DataLinkLayer) handleRetFrame(receivedFrame Frame) error {
	// 1. находим (как?) последний кадр
	// можно хранить массив последних кадров в DataLinkLayer
	// 2. отправляем еще раз
	return nil
}

func (d *DataLinkLayer) handleStateFrame(receivedFrame Frame) error {
	// 1. сохранить себе таблицу пользователей
	d.state = d.getStateFromFrame(receivedFrame)

	err := d.sendFrame(receivedFrame)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// 2. отправить запрос в интерфейс на разрешение действий
	return nil
}
