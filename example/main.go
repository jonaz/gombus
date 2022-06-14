package main

import (
	"fmt"

	"github.com/jonaz/gombus"
	"github.com/sirupsen/logrus"
)

func main() {
	// gombus.SendUD2(nil)

	conn, err := gombus.Dial("192.168.13.42:10001")
	if err != nil {
		logrus.Error(err)
		return
	}
	defer conn.Close()

	// frame := gombus.SetPrimaryUsingPrimary(0, 3)
	frame := gombus.RequestUD2(1)
	fmt.Printf("sending: % x\n", frame)
	n, err := conn.Write(frame)
	if err != nil {
		logrus.Error(err)
		return
	}

	logrus.Info("wrote n: ", n)
	resp, err := gombus.ReadLongFrame(conn)
	if err != nil {
		logrus.Error(err)
		return
	}

	fmt.Printf("read: % x\n", resp)
}
