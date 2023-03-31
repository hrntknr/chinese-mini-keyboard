package main

import "github.com/pkg/errors"

func getButtonCode(button string) (byte, error) {
	var key byte
	switch button {
	case "KEY1":
		key = 0x01
	case "KEY2":
		key = 0x02
	case "KEY3":
		key = 0x03
	case "KEY4":
		key = 0x04
	case "KEY5":
		key = 0x05
	case "KEY6":
		key = 0x06
	case "KEY7":
		key = 0x07
	case "KEY8":
		key = 0x08
	case "KEY9":
		key = 0x09
	case "KEY10":
		key = 0x0a
	case "KEY11":
		key = 0x0b
	case "KEY12":
		key = 0x0c
	case "K1_LEFT":
		key = 0x0d
	case "K1_RIGHT":
		key = 0x0e
	case "K2_LEFT":
		key = 0x0f
	case "K2_RIGHT":
		key = 0x10
	case "K3_LEFT":
		key = 0x11
	case "K3_RIGHT":
		key = 0x12
	default:
		return 0, errors.New("invalid button")
	}
	return key, nil
}
