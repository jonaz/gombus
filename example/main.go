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

	// logrus.SetLevel(logrus.DebugLevel)

	conn, err := gombus.Dial("192.168.13.42:10001")
	if err != nil {
		logrus.Error(err)
		return
	}
	defer conn.Close()

	frame := gombus.SndNKE(uint8(*primaryID))
	fmt.Printf("sending nke: % x\n", frame)
	_, err = conn.Write(frame)
	if err != nil {
		logrus.Error(err)
		return
	}
	_, err = gombus.ReadSingleCharFrame(conn)
	if err != nil {
		logrus.Error(err)
		return
	}

	// frame := gombus.SetPrimaryUsingPrimary(0, 3)
	respFrame := &gombus.DecodedFrame{}
	lastFCB := true
	frames := 0
	for respFrame.HasMoreRecords() || frames == 0 {
		frames++
		// frame := gombus.SetPrimaryUsingPrimary(0, 3)
		frame = gombus.RequestUD2(uint8(*primaryID))
		if !lastFCB {
			frame.SetFCB()
			frame.SetChecksum()
		}
		lastFCB = frame.C().FCB()

		fmt.Printf("sending: % x\n", frame)
		fmt.Printf("sending C: % b\n", frame.C())
		_, err = conn.Write(frame)
		if err != nil {
			logrus.Error(err)
			return
		}

		resp, err := conn.ReadLongFrame()
		if err != nil {
			logrus.Error(err)
			return
		}

		fmt.Printf("read: % x\n", resp)
		fmt.Printf("read C: % b\n", resp.C())

		respFrame, err = resp.Decode()
		if err != nil {
			logrus.Error(err)
			return
		}
		spew.Dump(respFrame)
		logrus.Info("read total values: ", len(respFrame.DataRecords))
	}

	logrus.Info("read total frames: ", frames)
}
