package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/jonaz/gombus"
)

var device = flag.String("device", "192.168.13.42:10001", "address. ether tcp ex 192.168.1.10:10000 or /dev/tty*")

var mode = flag.String("mode", "scan", "valid modes are: scan,set-primary")

// set-primary
var newPrimary = flag.Int("new-primary", -1, "new primary address")
var currentPrimary = flag.Int("current-primary", -1, "current primary address")

// scan
var addrStart = flag.Int("addr-start", 1, "primary address start number. use 254 to scan if only one meter connected")
var addrStop = flag.Int("addr-stop", 250, "primary address stop number")

func main() {
	flag.Parse()

	log.Printf("connecting to: %s\n", *device)
	conn, err := gombus.Dial(*device)
	if err != nil {
		log.Println(fmt.Errorf("error connecting to mbus: %w", err))
		return
	}

	if *mode == "scan" {
		for i := *addrStart; i <= *addrStop; i++ {
			err := readPrimaryAddress(conn, i)
			if err != nil {
				log.Println(fmt.Errorf("error connecting to mbus: %w", err))
			}
		}
		return
	}

	if *mode == "set-primary" {
		//TODO
		log.Printf("change primary address from %d to %d\n", *currentPrimary, *newPrimary)
	}
}
func readPrimaryAddress(conn *gombus.Conn, primaryAddr int) error {
	_, err := conn.Write(gombus.SndNKE(uint8(primaryAddr)))
	if err != nil {
		return err
	}
	err = conn.SetReadDeadline(time.Now().Add(1 * time.Second))
	if err != nil {
		return err
	}
	_, err = gombus.ReadSingleCharFrame(conn)
	if err != nil {
		return err
	}

	frame, err := gombus.ReadSingleFrame(conn, primaryAddr)
	if err != nil {
		return err
	}

	log.Println("Found device:", frame.SerialNumber, frame.Manufacturer, frame.Version, frame.DeviceType, frame.SecondaryAddressString())

	return nil
}
