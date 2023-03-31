package main

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

func parseKeys(keys []string) ([][]byte, error) {
	var buf [][]byte
	for _, k := range keys {
		b1, b2, err := parseKey(k)
		if err != nil {
			return nil, err
		}
		buf = append(buf, []byte{b1, b2})
	}
	return buf, nil
}

func parseKey(key string) (byte, byte, error) {
	var buf1, buf2 byte

	keys := strings.Split(key, "+")
	if len(keys) != 1 {
		for _, k := range keys[:1] {
			switch k {
			case "ctrl":
				buf1 |= 0x01
			case "shift":
				buf1 |= 0x02
			case "alt":
				buf1 |= 0x04
			case "gui":
				buf1 |= 0x08
			default:
				return 0, 0, errors.Errorf("unknown key: %s", k)
			}
		}
	}

	switch keys[len(keys)-1] {
	case "a":
		buf2 = 0x04
	case "b":
		buf2 = 0x05
	case "c":
		buf2 = 0x06
	case "d":
		buf2 = 0x07
	case "e":
		buf2 = 0x08
	case "f":
		buf2 = 0x09
	case "g":
		buf2 = 0x0a
	case "h":
		buf2 = 0x0b
	case "i":
		buf2 = 0x0c
	case "j":
		buf2 = 0x0d
	case "k":
		buf2 = 0x0e
	case "l":
		buf2 = 0x0f
	case "m":
		buf2 = 0x10
	case "n":
		buf2 = 0x11
	case "o":
		buf2 = 0x12
	case "p":
		buf2 = 0x13
	case "q":
		buf2 = 0x14
	case "r":
		buf2 = 0x15
	case "s":
		buf2 = 0x16
	case "t":
		buf2 = 0x17
	case "u":
		buf2 = 0x18
	case "v":
		buf2 = 0x19
	case "w":
		buf2 = 0x1a
	case "x":
		buf2 = 0x1b
	case "y":
		buf2 = 0x1c
	case "z":
		buf2 = 0x1d
	case "1":
		buf2 = 0x1e
	case "2":
		buf2 = 0x1f
	case "3":
		buf2 = 0x20
	case "4":
		buf2 = 0x21
	case "5":
		buf2 = 0x22
	case "6":
		buf2 = 0x23
	case "7":
		buf2 = 0x24
	case "8":
		buf2 = 0x25
	case "9":
		buf2 = 0x26
	case "0":
		buf2 = 0x27
	case "enter":
		buf2 = 0x28
	case "esc":
		buf2 = 0x29
	case "backspace":
		buf2 = 0x2a
	case "tab":
		buf2 = 0x2b
	case "space":
		buf2 = 0x2c
	case "-":
		buf2 = 0x2d
	case "=":
		buf2 = 0x2e
	case "[":
		buf2 = 0x2f
	case "]":
		buf2 = 0x30
	case "\\":
		buf2 = 0x31
	case "#":
		buf2 = 0x32
	case ";":
		buf2 = 0x33
	case "'":
		buf2 = 0x34
	case "`":
		buf2 = 0x35
	case ",":
		buf2 = 0x36
	case ".":
		buf2 = 0x37
	case "/":
		buf2 = 0x38
	case "capslock":
		buf2 = 0x39
	case "f1":
		buf2 = 0x3a
	case "f2":
		buf2 = 0x3b
	case "f3":
		buf2 = 0x3c
	case "f4":
		buf2 = 0x3d
	case "f5":
		buf2 = 0x3e
	case "f6":
		buf2 = 0x3f
	case "f7":
		buf2 = 0x40
	case "f8":
		buf2 = 0x41
	case "f9":
		buf2 = 0x42
	case "f10":
		buf2 = 0x43
	case "f11":
		buf2 = 0x44
	case "f12":
		buf2 = 0x45
	case "printscreen":
		buf2 = 0x46
	case "scrolllock":
		buf2 = 0x47
	case "pause":
		buf2 = 0x48
	case "insert":
		buf2 = 0x49
	case "home":
		buf2 = 0x4a
	case "pageup":
		buf2 = 0x4b
	case "delete":
		buf2 = 0x4c
	case "end":
		buf2 = 0x4d
	case "pagedown":
		buf2 = 0x4e
	case "right":
		buf2 = 0x4f
	case "left":
		buf2 = 0x50
	case "down":
		buf2 = 0x51
	case "up":
		buf2 = 0x52
	case "numlock":
		buf2 = 0x53
	case "kp-/":
		buf2 = 0x54
	case "kp-*":
		buf2 = 0x55
	case "kp-+":
		buf2 = 0x56
	case "kp-enter":
		buf2 = 0x58
	case "kp-1":
		buf2 = 0x59
	case "kp-2":
		buf2 = 0x5a
	case "kp-3":
		buf2 = 0x5b
	case "kp-4":
		buf2 = 0x5c
	case "kp-5":
		buf2 = 0x5d
	case "kp-6":
		buf2 = 0x5e
	case "kp-7":
		buf2 = 0x5f
	case "kp-8":
		buf2 = 0x60
	case "kp-9":
		buf2 = 0x61
	case "kp-0":
		buf2 = 0x62
	case "kp-.":
		buf2 = 0x63
	default:
		return 0, 0, fmt.Errorf("unknown key: %s", key)
	}

	return buf1, buf2, nil
}

func getMediaKey(key string) (byte, error) {
	// layerが違うとコードも違うみたいなのでlayer supportを追加する時は注意
	var code byte
	switch key {
	case "play":
		code = 0xcd
	case "prev":
		code = 0xb6
	case "next":
		code = 0xb5
	case "mute":
		code = 0xe2
	case "volup":
		code = 0xe9
	case "voldown":
		code = 0xea
	default:
		return 0, fmt.Errorf("unknown media key: %s", key)
	}
	return code, nil
}
