package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/google/gousb"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var log *zap.Logger

func init() {
	cfg := zap.NewDevelopmentConfig()
	switch strings.ToLower(os.Getenv("LOG_LEVEL")) {
	case "debug":
		cfg.Level.SetLevel(zap.DebugLevel)
	case "info":
		cfg.Level.SetLevel(zap.InfoLevel)
	case "warn":
		cfg.Level.SetLevel(zap.WarnLevel)
	case "error":
		cfg.Level.SetLevel(zap.ErrorLevel)
	case "dpanic":
		cfg.Level.SetLevel(zap.DPanicLevel)
	case "panic":
		cfg.Level.SetLevel(zap.PanicLevel)
	}
	log, _ = cfg.Build()
}

func main() {
	var rootCmd = &cobra.Command{Use: "app"}
	var key = &cobra.Command{
		Use:   "key",
		Short: "set key to button",
		Run: func(cmd *cobra.Command, args []string) {
			if err := setKey(args); err != nil {
				log.Fatal(fmt.Sprintf("%+v", err))
			}
		},
	}
	var mediakey = &cobra.Command{
		Use:   "mediakey",
		Short: "set media key to button",
		Run: func(cmd *cobra.Command, args []string) {
			if err := setMediaKey(args); err != nil {
				log.Fatal(fmt.Sprintf("%+v", err))
			}
		},
	}
	var led = &cobra.Command{
		Use:   "led",
		Short: "set led mode",
		Run: func(cmd *cobra.Command, args []string) {
			if err := setLed(args); err != nil {
				log.Fatal(fmt.Sprintf("%+v", err))
			}
		},
	}
	rootCmd.AddCommand(key)
	rootCmd.AddCommand(mediakey)
	rootCmd.AddCommand(led)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(fmt.Sprintf("%+v", err))
	}
}

func setKey(args []string) error {
	if len(args) < 2 {
		return errors.New("invalid args")
	}
	button, err := getButtonCode(args[0])
	if err != nil {
		return errors.Wrap(err, "failed to get button code")
	}
	keys, err := parseKeys(args[1:])
	if err != nil {
		return errors.Wrap(err, "failed to parse key")
	}
	if len(keys) > 5 {
		return errors.New("too many keys")
	}
	for i, key := range keys {
		if i != 0 && key[0] != 0x00 {
			return errors.New("modifier is only available for the first key")
		}
	}
	length := byte(len(keys))

	ep, err := initUsb(0x1189, 0x8890)
	if err != nil {
		return errors.Wrap(err, "failed to init usb")
	}

	if err := sendCode(ep, []byte{0x03, 0xa1, 0x01}); err != nil {
		return errors.Wrap(err, "failed to send code")
	}
	for i, key := range keys {
		if i == 0 {
			if err := sendCode(ep, []byte{0x03, button, 0x11, length, 0x00, key[0]}); err != nil {
				return errors.Wrap(err, "failed to send code")
			}
		}
		if err := sendCode(ep, []byte{0x03, button, 0x11, length, byte(i + 1), key[0], key[1]}); err != nil {
			return errors.Wrap(err, "failed to send code")
		}
	}
	if err := sendCode(ep, []byte{0x03, 0xaa, 0xaa}); err != nil {
		return errors.Wrap(err, "failed to send code")
	}
	return nil
}

func setMediaKey(args []string) error {
	if len(args) != 2 {
		return errors.New("invalid args")
	}
	button, err := getButtonCode(args[0])
	if err != nil {
		return errors.Wrap(err, "failed to get button code")
	}
	key, err := getMediaKey(args[1])
	if err != nil {
		return errors.Wrap(err, "failed to get media key")
	}

	ep, err := initUsb(0x1189, 0x8890)
	if err != nil {
		return errors.Wrap(err, "failed to init usb")
	}

	if err := sendCode(ep, []byte{0x03, 0xa1, 0x01}); err != nil {
		return errors.Wrap(err, "failed to send code")
	}
	if err := sendCode(ep, []byte{0x03, button, 0x12, key}); err != nil {
		return errors.Wrap(err, "failed to send code")
	}
	if err := sendCode(ep, []byte{0x03, 0xaa, 0xaa}); err != nil {
		return errors.Wrap(err, "failed to send code")
	}

	return nil
}

func setLed(args []string) error {
	if len(args) != 1 {
		return errors.New("invalid args")
	}

	var mode byte
	switch args[0] {
	case "MODE0":
		mode = 0
	case "MODE1":
		mode = 1
	case "MODE2":
		mode = 2
	default:
		return errors.New("invalid mode")
	}

	ep, err := initUsb(0x1189, 0x8890)
	if err != nil {
		return errors.Wrap(err, "failed to init usb")
	}

	if err := sendCode(ep, []byte{0x03, 0xa1, 0x01}); err != nil {
		return errors.Wrap(err, "failed to send code")
	}
	if err := sendCode(ep, []byte{0x03, 0xb0, 0x18, mode}); err != nil {
		return errors.Wrap(err, "failed to send code")
	}
	if err := sendCode(ep, []byte{0x03, 0xaa, 0xa1}); err != nil {
		return errors.Wrap(err, "failed to send code")
	}

	return nil
}

func sendCode(ep *gousb.OutEndpoint, payload []byte) error {
	log.Debug("send code", zap.String("payload", hex.EncodeToString(payload)))
	buf := make([]byte, 65)
	copy(buf, payload)
	if _, err := ep.Write(buf); err != nil {
		return errors.Wrap(err, "failed to write")
	}
	return nil
}

func initUsb(vid, pid gousb.ID) (*gousb.OutEndpoint, error) {
	ctx := gousb.NewContext()
	dev, err := ctx.OpenDeviceWithVIDPID(vid, pid)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open device")
	}
	defer dev.Close()

	cfg, err := dev.Config(1)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open config")
	}
	iface, err := cfg.Interface(1, 0)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open interface")
	}
	ep, err := iface.OutEndpoint(2)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open endpoint")
	}

	return ep, nil
}
