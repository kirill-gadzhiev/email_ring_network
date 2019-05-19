package main

import (
	"bytes"
	"errors"
)

type FrameParser struct {
	buff []byte
}

func (f *FrameParser) Empty() bool {
	err := f.trim()
	if err != nil {
		return true
	}

	if len(f.buff) > 0 {
		if bytes.Count(f.buff, []byte{START_BYTE}) > 1 {
			return false
		}
	}
	return true
}

func (f *FrameParser) Top() (byte, error) {
	if len(f.buff) > 0 {
		return f.buff[0], nil
	}
	return 0, errors.New("empty buffer")
}

func (f *FrameParser) Pop() (byte, error) {
	if len(f.buff) > 0 {
		temp := f.buff[0]
		f.buff = f.buff[1:]
		return temp, nil
	}
	return 0,  errors.New("empty buffer")
}

func (f *FrameParser) trim() error {
	for {
		value, err := f.Top()
		if err != nil {
			return errors.New("empty buffer")
		}
		if value == START_BYTE {
			return nil
		}
		_, err = f.Pop()
		if err != nil {
			return err
		}
	}
}

func (f *FrameParser) PopFrame() ([]byte, error) {
	if bytes.Count(f.buff, []byte{START_BYTE}) < 2 {
		return nil, errors.New("no complete frames in buffer")
	}

	err := f.trim()
	if err != nil {
		return nil, err
	}

	firstByte, _ := f.Pop()
	frame := []byte{firstByte}
	for {
		nextByte, err := f.Pop()
		if err != nil {
			return nil, err
		}

		frame = append(frame, nextByte)
		if nextByte == STOP_BYTE {
			return frame, nil
		}
	}
}

func (f *FrameParser) AddBytes(data []byte) {
	f.buff = append(f.buff, data...)
}