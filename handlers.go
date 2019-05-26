package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Equanox/gotron"
	"go.bug.st/serial.v1"
)

func master(backPortMode *serial.Mode, backPortName string, frontPortMode *serial.Mode, frontPortName string) error {
	// 0. Подключиться к заднему и переднему портам
	fmt.Println("try to open backport")
	var backPort DataLinkLayer
	err := backPort.logicalConnect(backPortMode, backPortName)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("try to open frontport")
	var frontPort DataLinkLayer
	err = frontPort.logicalConnect(frontPortMode, frontPortName)
	if err != nil {
		fmt.Println(err)
		return err
	}

	core2interface := make(chan SomeEvent)
	interface2core := make(chan SomeEvent)
	fmt.Println("chans made")


	go initializeInterface(core2interface, interface2core, []string{}, &frontPort)
	fmt.Println("interface made")

	emailEvent := <-interface2core
	fmt.Println("got event from chan ", emailEvent.GetEvent(), emailEvent)
	if emailEvent.GetEvent() != SET_USER {
		fmt.Println("expected SET_USER event, got:", emailEvent.GetEvent())
		return errors.New("wrong event type")
	}

	// 1. Отправить LINK_FRAME
	fmt.Println("cast email")
	userCasted, ok := emailEvent.(UserEvent)
	fmt.Println("OK: ", ok, "userCasted: ", userCasted)
	if !ok {
		return errors.New("wrong event type")
	}

	fmt.Println("sendLinkFrameByMaster", userCasted.User.Email)
	err = frontPort.sendLinkFrameByMaster(userCasted.User.Email)
	fmt.Println("send link frame")
	if err != nil {
		fmt.Println(err)
		return err
	}

	parser := &FrameParser{[]byte{}}
	fmt.Println("frameparser made")

	// 2. Дождаться получения LINK_FRAME после круга по кольцу
	ch := make(chan []byte)
	go backPort.listen(ch)
	fmt.Println("start listening")

	exit := false
	var linkFrame *Frame
	for dataPart := range ch {
		parser.AddBytes(dataPart)
		for {
			if parser.Empty() {
				break
			}

			frameBytes, err := parser.PopFrame()
			if err != nil {
				fmt.Println(err)
				break
			}

			linkFrame, err = parseBytesToFrame(frameBytes)
			fmt.Println("got link frame after ring: ", string(linkFrame.data))

			if err != nil {
				fmt.Println(err)
			}

			if linkFrame.frameType != LINK_FRAME {
				// тут по хорошему надо делать что-то
				// ситуация, если кто-то прислал кадр то инициализации кольца
			}

			exit = true
			break
		}

		if exit {
			break
		}
	}

	frontPort.state = frontPort.getStateFromFrame(*linkFrame)

	fmt.Println("state: ", frontPort.state)

	// 3. Отправить STATE_FRAME
	err = frontPort.sendStateFrame(linkFrame.data, MIN_ADDRESS)
	fmt.Println("send state frame")
	if err != nil {
		fmt.Println(err)
	}

	// 4. Дождаться получения STATE_FRAME
	exit = false
	var stateFrame *Frame
	for dataPart := range ch {
		parser.AddBytes(dataPart)
		for {
			if parser.Empty() {
				break
			}

			frameBytes, err := parser.PopFrame()
			if err != nil {
				fmt.Println(err)
				break
			}

			stateFrame, err = parseBytesToFrame(frameBytes)
			fmt.Println("got state frame after ring: ", string(stateFrame.data))

			if err != nil {
				fmt.Println(err)
			}

			if stateFrame.frameType != STATE_FRAME {
				// тут по хорошему надо делать что-то
				// ситуация, если кто-то прислал кадр то инициализации кольца
			}

			exit = true
			break
		}

		if exit {
			break
		}
	}

	// передача стейта в горутину с интерфейсом
	availableUsers := []User{}
	for _, email := range frontPort.state {
		availableUsers = append(availableUsers, User{email})
	}
	status := NetworkStatus{Connection:true, AvailableUsers: availableUsers}
	stateEvent := NetworkStatusEvent{&gotron.Event{Event:NETWORK_STATUS}, status}
	core2interface <- stateEvent

	// 5. Стать Slave
	fmt.Println("now master = slave")

	for dataPart := range ch {
		parser.AddBytes(dataPart)
		for {
			if parser.Empty() {
				break
			}

			frameBytes, err := parser.PopFrame()
			if err != nil {
				fmt.Println(err)
				break
			}
			frame, err := parseBytesToFrame(frameBytes)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("Data: ", string(frame.data))
			err = frontPort.handleFrame(*frame, core2interface, interface2core)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}

func slave(backPortMode *serial.Mode, backPortName string, frontPortMode *serial.Mode, frontPortName string) error {
	// 0. Подключиться к заднему и переднему портам
	var backPort DataLinkLayer
	fmt.Println("TRY TO OPEN BACKPORT: ", backPortName)
	err := backPort.logicalConnect(backPortMode, backPortName)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("TRY TO OPEN FRONTPORT: ", frontPortName)
	var frontPort DataLinkLayer
	err = frontPort.logicalConnect(frontPortMode, frontPortName)
	if err != nil {
		fmt.Println(err)
		return err
	}

	core2interface := make(chan SomeEvent)
	interface2core := make(chan SomeEvent)
	fmt.Println("chans made")

	parser := &FrameParser{[]byte{}}
	fmt.Println("parser made")
	// 1. Слушаем задний
	ch := make(chan []byte, 3)
	fmt.Println("start listen")
	go backPort.listen(ch)
	fmt.Println("after go listen")
	// 2. Парсим фрейм
	// 3. switch-case по frameType
	for dataPart := range ch {
		fmt.Println("range ch")
		parser.AddBytes(dataPart)
		for {
			if parser.Empty() {
				break
			}

			frameBytes, err := parser.PopFrame()
			if err != nil {
				fmt.Println(err)
				break
			}


			frame, err := parseBytesToFrame(frameBytes)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("Data: ", string(frame.data))
			err = frontPort.handleFrame(*frame, core2interface, interface2core)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}


func (d *DataLinkLayer) handleFrame(receivedFrame Frame, core2interface chan SomeEvent, interface2core chan SomeEvent) error {

	// здесь мы пересылаем кадр если он не предназначен нам
	if receivedFrame.to != d.currentAddress {
		if receivedFrame.to != BROADCAST {
			return d.sendFrame(receivedFrame)
		}
	}

	switch receivedFrame.frameType {
	case INFO_FRAME:
		fmt.Println("INFO_FRAME")
		return d.handleInfoFrame(receivedFrame, core2interface)
	case LINK_FRAME:
		fmt.Println("LINK_FRAME")
		// это только для slave компа
		return d.handleLinkFrame(receivedFrame, core2interface, interface2core)
	case UPLINK_FRAME:
		fmt.Println("UPLINK_FRAME")
		return d.handleUplinkFrame(receivedFrame, core2interface)
	case ACK_FRAME:
		fmt.Println("ACK_FRAME")
		return d.handleAckFrame(receivedFrame)
	case RET_FRAME:
		fmt.Println("RET_FRAME")
		return d.handleRetFrame(receivedFrame)
	case STATE_FRAME:
		fmt.Println("STATE_FRAME")
		return d.handleStateFrame(receivedFrame, core2interface)
	default:

		return errors.New("unknown type of Frame")
	}
	return nil
}

func (d *DataLinkLayer) handleInfoFrame(receivedFrame Frame, core2interface chan SomeEvent) error {
	// 1. обработать фрейм

	var letter Letter
	err := json.Unmarshal(receivedFrame.data, &letter)
	if err != nil {
		fmt.Println(err)
		return err
	}
	messageEvent := MessageReceivedEvent{&gotron.Event{Event:MESSAGE_RECEIVED}, letter}
	fmt.Println("send message event to interface", messageEvent)
	core2interface <- messageEvent
	// отправляем в горутину с интерфейсом инфу

	// 2. отправить ack
	err = d.sendAckFrame(receivedFrame.from, receivedFrame.to)
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

func (d *DataLinkLayer) handleLinkFrame(receivedFrame Frame, core2interface chan SomeEvent, interface2core chan SomeEvent) error {
	availableUsersMap := d.getStateFromFrame(receivedFrame)
	availableUsers := []string{}

	for  _, value := range availableUsersMap {
		availableUsers = append(availableUsers, value)
	}

	go initializeInterface(core2interface, interface2core, availableUsers, d)

	emailEvent := <-interface2core
	if emailEvent.GetEvent() == SET_USER {
		fmt.Println("cast email")
		userCasted, ok := emailEvent.(UserEvent)
		fmt.Println("OK: ", ok, "userCasted: ", userCasted)
		if !ok {
			return errors.New("wrong event type")
		}
		return d.sendLinkFrameBySlave(receivedFrame, userCasted.User.Email)
	}

	// 2. отправить link далее

	return errors.New("wrong event from interface")
}

func (d *DataLinkLayer) handleUplinkFrame(receivedFrame Frame, core2interface chan SomeEvent) error {
	// 1. отправляем uplink далее
	if receivedFrame.from != d.currentAddress {
		err := d.sendFrame(receivedFrame)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	// 2. разрываем соединение
	err := d.logicalDisconnect()
	if err != nil {
		return err
	}

	status := NetworkStatus{Connection:false, AvailableUsers: []User{}}
	stateEvent := NetworkStatusEvent{&gotron.Event{Event:NETWORK_STATUS}, status}
	core2interface <- stateEvent

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

func (d *DataLinkLayer) handlePingFrame(receivedFrame Frame) error {
	// 1. выставляем что соединение есть
	// 2. отправляем новый ping frame
	// 3. если не приходит за PING_TIMEOUT то выставляем отсутвие соединения
	// возвращаем
	return nil
}

func (d *DataLinkLayer) handleStateFrame(receivedFrame Frame, core2interface chan SomeEvent) error {
	// 1. сохранить себе таблицу пользователей
	d.state = d.getStateFromFrame(receivedFrame)

	fmt.Println("now will send state frame")
	err := d.sendFrame(receivedFrame)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// 2. отправить запрос в интерфейс на разрешение действий
	availableUsers := []User{}
	for _, email := range d.state {
		availableUsers = append(availableUsers, User{email})
	}
	status := NetworkStatus{Connection:true, AvailableUsers: availableUsers}
	stateEvent := NetworkStatusEvent{&gotron.Event{Event:NETWORK_STATUS}, status}
	core2interface <- stateEvent

	return nil
}
