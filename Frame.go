package main

import (
	"errors"
	"log"
)

// формат кадра (* - необязательные поля)
//    0     1     2     3     4     5     ?
//  start   to  from  type  d_len data stop
// [ 1 ] [ 1 ] [ 1 ] [ 1 ] [ 1*] [ N*] [ 1 ]

// таким образом служебные кадры без данных имеют длину 5 байт

// для LINK_FRAME в поле data данные о клиентах сети должны располагаться в следующем формате:
// [addr] [email]\n
// где addr от MIN_ADDRESS до MAX_ADDRESS
//     email в виде abc123@some.kek
// addr и email разделяются пробелом, записи о клиентах разделяются переносом строки \n

const (
	INFO_FRAME   = 0x81  // для передачи коротких сообщений
	LINK_FRAME   = 0x82  // для установки логического соединения (в нашем случае "Кольцо")
	UPLINK_FRAME = 0x83  // для разрыва логич. соед
	ACK_FRAME    = 0x84  // для подтвержения получения последнего кадра
	RET_FRAME    = 0x85  // для запроса на повторение отправки последнего кадра

	STATE_FRAME  = 0x86  // для передачи сведений о всех клиентах сети
	PING_FRAME   = 0x87
)

const (
	START_BYTE = 0xFF
	STOP_BYTE  = 0xFF
)

const (
	MIN_ADDRESS  = 0x01
	MAX_ADDRESS  = 0x7E
	BROADCAST    = 0x7F
)

type Frame struct {
	frameType     byte
	to            byte
	from          byte
	dataLength    byte
	data        []byte
}

func createFrame(frameType byte, to byte, from byte, data []byte) (*Frame, error) {
	switch frameType {
	case INFO_FRAME: fallthrough
	case LINK_FRAME:
		return &Frame{frameType, to, from, byte(len(data)), data}, nil
	case UPLINK_FRAME: fallthrough
	case ACK_FRAME: fallthrough
	case RET_FRAME:
		return &Frame{frameType, to, from, 0, nil}, nil
	default:
		return nil, errors.New("unknown type of Frame")
	}
}

func (f *Frame) toBytesArray() []byte {
	arr := []byte{START_BYTE, f.to, f.from, f.frameType}

	if f.dataLength != 0 && f.data != nil {
		arr = append(arr, f.dataLength)
		arr = append(arr, f.data...)
	}

	arr = append(arr, STOP_BYTE)
	return arr
}

func parseBytesToFrame(bytes []byte) (*Frame, error) {
	err := validateBytes(bytes)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	to := bytes[1]
	from := bytes[2]
	frameType := bytes[3]

	switch bytes[3] {
	case INFO_FRAME: fallthrough
	case LINK_FRAME:
		n := len(bytes)
		data := bytes[5:n-1]
		return &Frame{frameType, to, from, byte(len(data)), data}, nil
	case UPLINK_FRAME: fallthrough
	case ACK_FRAME: fallthrough
	case RET_FRAME:
		return &Frame{frameType, to, from, 0, nil}, nil
	default:
		return nil, errors.New("unknown type of Frame")
	}
}

func validateBytes(bytes []byte) error {
	n := len(bytes)

	if bytes[0] != START_BYTE {
		return errors.New("invalid start byte")
	}

	if bytes[n-1] != STOP_BYTE {
		return errors.New("invalid stop byte")
	}

	to := bytes[1]
	from := bytes[2]

	if to < MIN_ADDRESS || to > BROADCAST {
		return errors.New("\"TO\" field out of address range")
	}

	if from < MIN_ADDRESS || from > MAX_ADDRESS {
		return errors.New("\"FROM\" field out of address range")
	}

	frameType := bytes[3]
	switch frameType {
	case INFO_FRAME: fallthrough
	case LINK_FRAME: fallthrough
	case UPLINK_FRAME: fallthrough
	case ACK_FRAME: fallthrough
	case RET_FRAME:
		return nil
	default:
		return errors.New("unknown type of Frame")
	}

	dataLength := bytes[4]
	if dataLength != STOP_BYTE { // если есть поле длины данных
		if dataLength == 0 { // если оно равно 0
			if bytes[5] != STOP_BYTE { // то след байт должен быть STOP_BYTE
				return errors.New("invalid \"dataLength\" field: there are data, but length = 0") // иначе ошибка
			}
		} else { // если поле длины не равно 0, то проверяем, что такое кол-во байт есть в поле данных
			expectedFrameLength := 4 + int(dataLength)
			if expectedFrameLength > n {
				return errors.New("invalid \"dataLength\" field") // иначе ошибка
			}
		}
	}
	return nil
}