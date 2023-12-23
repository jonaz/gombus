package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/jonaz/gombus"
)

var device = flag.String("device", "192.168.13.42:10001", "address. ether tcp ex 192.168.1.10:10000 or /dev/tty*")

var mode = flag.String("mode", "scan", "valid modes are: scan,set-primary,read-single,read-multi")

// set-primary
var newPrimary = flag.Int("new-primary", -1, "new primary address")
var currentPrimary = flag.Int("current-primary", -1, "current primary address")

// scan
var addr = flag.Int("addr", 1, "primary address start number. use 254 if only one meter connected to bus")
var addrStop = flag.Int("addr-stop", 250, "primary address stop number")

func main() {
	flag.Parse()

	log.Printf("connecting to: %s\n", *device)
	conn, err := dial(*device)
	if err != nil {
		log.Println(fmt.Errorf("error connecting to mbus: %w", err))
		return
	}

	if *mode == "scan" {
		for i := *addr; i <= *addrStop; i++ {
			frame, err := readPrimaryAddress(conn, i)
			if err != nil {
				log.Println(err)
				continue
			}
			log.Println("Found device:", frame.SerialNumber, frame.Manufacturer, frame.Version, frame.DeviceType, frame.SecondaryAddressString())
		}
		return
	}

	if *mode == "set-primary" {
		//TODO
		log.Printf("change primary address from %d to %d\n", *currentPrimary, *newPrimary)
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
	_, err := conn.Write(gombus.SndNKE(uint8(primaryAddr)))
	if err != nil {
		return nil, err
	}
	err = conn.SetReadDeadline(time.Now().Add(1 * time.Second))
	if err != nil {
		return nil, err
	}
	_, err = gombus.ReadSingleCharFrame(conn)
	if err != nil {
		return nil, err
	}

	frame, err := gombus.ReadSingleFrame(conn, primaryAddr)
	if err != nil {
		return nil, err
	}

	return frame, nil
}
