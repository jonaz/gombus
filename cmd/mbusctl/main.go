package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/jonaz/gombus"
	"github.com/sirupsen/logrus"
)

var device = flag.String("device", "192.168.13.42:10001", "address. ether tcp ex 192.168.1.10:10000 or /dev/tty*")

var mode = flag.String("mode", "scan", "valid modes are: scan,set-primary,read-single,read-multi")

// set-primary
var newPrimary = flag.Int("new-primary", -1, "new primary address")
var currentPrimary = flag.Int("current-primary", -1, "current primary address")

// scan
var addr = flag.Int("addr", 1, "primary address start number. use 254 if only one meter connected to bus")
var addrStop = flag.Int("addr-stop", 250, "primary address stop number")
var logLevel = flag.String("loglevel", "info", "available levels are: "+strings.Join(getLevels(), ","))

func getLevels() []string {
	lvls := make([]string, len(logrus.AllLevels))
	for k, v := range logrus.AllLevels {
		lvls[k] = v.String()
	}
	return lvls
}

func main() {
	flag.Parse()
	lvl, err := logrus.ParseLevel(*logLevel)
	if err != nil {
		log.Fatal(err)
	}
	if lvl != logrus.InfoLevel {
		fmt.Fprintf(os.Stderr, "using loglevel: %s\n", lvl.String())
	}
	logrus.SetLevel(lvl)

	log.Printf("connecting to: %s\n", *device)
	conn, err := dial(*device)
	if err != nil {
		log.Println(fmt.Errorf("error connecting to mbus: %w", err))
		return
	}

	if *mode == "scan" {
		for i := *addr; i <= *addrStop; i++ {
			log.Println("checking address:", i)
			frame, err := readPrimaryAddress(conn, i)
			if err != nil {
				log.Println("error checking address:", i, err)
				continue
			}
			log.Println("Found device:", frame.SerialNumber, frame.Manufacturer, frame.Version, frame.DeviceType, frame.SecondaryAddressString())
		}
		return
	}

	if *mode == "set-primary" {
		log.Printf("change primary address from %d to %d\n", *currentPrimary, *newPrimary)
		err := setPrimary(conn, *currentPrimary, *newPrimary)
		if err != nil {
			log.Println(err)
			return
		}
	}

	if *mode == "read-single" {
		frame, err := readPrimaryAddress(conn, *addr)
		if err != nil {
			log.Println(err)
			return
		}
		b, err := json.MarshalIndent(frame, "", "\t")
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(string(b))
		return
	}
	if *mode == "read-multi" {
		//TODO
	}
}

func dial(device string) (gombus.Conn, error) {
	_, _, err := net.SplitHostPort(device)
	if err != nil {
		log.Printf("device %s does not contain a port, assuming its a serial device\n", device)
		return gombus.DialSerial(device)
	}

	conn, err := gombus.DialTCP(device)
	if err != nil {
		return nil, fmt.Errorf("error connecting to mbus: %w", err)
	}

	return conn, nil
}

func readPrimaryAddress(conn gombus.Conn, primaryAddr int) (*gombus.DecodedFrame, error) {
	logrus.Trace("writing NKE")
	_, err := conn.Write(gombus.SndNKE(uint8(primaryAddr)))
	if err != nil {
		return nil, err
	}
	logrus.Trace("writing NKE done")

	logrus.Trace("conn.SetReadDeadline")
	err = conn.SetReadDeadline(time.Now().Add(1 * time.Second))
	if err != nil {
		return nil, err
	}

	logrus.Trace("conn.SetReadDeadline done")
	logrus.Trace("ReadSingleCharFrame")
	_, err = gombus.ReadSingleCharFrame(conn)
	if err != nil {
		return nil, err
	}
	logrus.Trace("ReadSingleCharFrame done")

	logrus.Trace("gombus.ReadSingleFrame")
	frame, err := gombus.ReadSingleFrame(conn, primaryAddr)
	if err != nil {
		return nil, err
	}
	logrus.Trace("gombus.ReadSingleFrame done")

	return frame, nil
}
func setPrimary(conn gombus.Conn, primaryAddr, newPrimary int) error {
	logrus.Trace("writing NKE")
	_, err := conn.Write(gombus.SndNKE(uint8(primaryAddr)))
	if err != nil {
		return err
	}
	logrus.Trace("writing NKE done")

	logrus.Trace("conn.SetReadDeadline")
	err = conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		return err
	}

	logrus.Trace("conn.SetReadDeadline done")
	logrus.Trace("ReadSingleCharFrame")
	_, err = gombus.ReadSingleCharFrame(conn)
	if err != nil {
		return err
	}
	logrus.Trace("ReadSingleCharFrame done")

	frame := gombus.SetPrimaryUsingPrimary(uint8(primaryAddr), uint8(newPrimary))
	_, err = conn.Write(frame)
	if err != nil {
		return err
	}
	_, err = gombus.ReadSingleCharFrame(conn)
	if err != nil {
		return fmt.Errorf("timeout waiting for answer after setting address: %w", err)
	}

	return nil
}
