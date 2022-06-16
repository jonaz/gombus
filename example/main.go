package main

import (
	"flag"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/jonaz/gombus"
	"github.com/sirupsen/logrus"
)

var primaryID = flag.Int("id", 1, "primaryID to fetch data from")

func main() {
	flag.Parse()

	conn, err := gombus.Dial("192.168.13.42:10001")
	if err != nil {
		logrus.Error(err)
		return
	}
	defer conn.Close()

	// frame := gombus.SetPrimaryUsingPrimary(0, 3)
	frame := gombus.RequestUD2(uint8(*primaryID))
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

	respFrame, err := resp.Decode()
	if err != nil {
		logrus.Error(err)
		return
	}

	spew.Dump(respFrame)
}
