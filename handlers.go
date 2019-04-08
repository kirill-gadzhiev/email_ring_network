package main

import (
	"errors"
	"go.bug.st/serial.v1"
)

func master(backPortMode serial.Mode, backPortName string, frontPortMode serial.Mode, frontPortName string) {
	// 0. Подключиться к заднему и переднему портам
	// 1. Отправить LINK_FRAME
	// 2. Дождаться получения LINK_FRAME после круга по кольцу
	// 3. Отправить STATE_FRAME
	// 4. Дождаться получения STATE_FRAME
	// 5. Стать Slave
}

func slave(backPortMode serial.Mode, backPortName string, frontPortMode serial.Mode, frontPortName string) {
	// 0. Подключиться к заднему и переднему портам
	// 1. Слушаем задний
	// 2. Парсим фрейм
	// 3. switch-case по frameType
}


func (d *DataLinkLayer) handleFrame(receivedFrame Frame) error {
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
	// 2. отправить ack
}

func (d *DataLinkLayer) handleLinkFrame(receivedFrame Frame) error {
	// 1. сохранить свой адрес
	// 2. отправить link далее
}

func (d *DataLinkLayer) handleUplinkFrame(receivedFrame Frame) error {
	// 1. отправляем uplink далее
	// 2. разрываем соединение
}

func (d *DataLinkLayer) handleAckFrame(receivedFrame Frame) error {
	// 1. просто отправляем (?)
}

func (d *DataLinkLayer) handleRetFrame(receivedFrame Frame) error {
	// 1. просто отправляем (?)
}

func (d *DataLinkLayer) handleStateFrame(receivedFrame Frame) error {
	// 1. сохранить себе таблицу пользователей
	// 2. отправить запрос в интерфейс на разрешение действий
}

