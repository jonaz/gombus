package gombus

import (
	"bufio"
	"fmt"
	"io"
	"time"
)

/*
// Does not handle timeout! the new one in tcp.go does
// perhaps use io.ReadCloser and close after timeout?

	func ReadLongFrame(r io.Reader) (LongFrame, error) {
		buf := bufio.NewReader(r)
		// TODO if this should be used again must implement whole frame detection from conn.ReadLongFrame
		msg, err := buf.ReadBytes(0x16)
		if err != nil {
			return nil, err
		}
		return LongFrame(msg), nil
	}
*/
var ErrNoLongFrameFound = fmt.Errorf("no long frame found")

func ReadLongFrame(conn Conn) (LongFrame, error) {
	buf := make([]byte, 4096)
	tmp := make([]byte, 4096)

	// foundStart := false
	length := 0
	globalN := -1
	for {
		err := conn.SetReadDeadline(time.Now().Add(2 * time.Second))
		if err != nil {
			return LongFrame{}, fmt.Errorf("error from SetReadDeadline: %w", err)
		}

		n, err := conn.Read(tmp)
		if err != nil {
			return LongFrame{}, fmt.Errorf("error reading from connection: %w", err)
		}

		for _, b := range tmp[:n] {
			globalN++
			buf[globalN] = b

			if globalN > 256 {
				return LongFrame{}, ErrNoLongFrameFound
			}

			// look for end byte after length +C+A+CI+checksum
			if length != 0 && globalN > length+4 && b == 0x16 {
				return LongFrame(buf[:globalN+1]), nil
			}

			// look for start sequence 68 LL LL 68
			if length == 0 && buf[0] == 0x68 && buf[3] == 0x68 && buf[1] == buf[2] {
				length = int(buf[1])
			}
		}
	}
}

func ReadSingleCharFrame(r io.Reader) (LongFrame, error) {
	buf := bufio.NewReader(r)
	msg, err := buf.ReadBytes(SingleCharacterFrame)
	if err != nil {
		return nil, err
	}
	return LongFrame(msg), nil
}

// ReadAnyAndPrint is used for debugging.
func ReadAnyAndPrint(r io.Reader) error {
	tmp := make([]byte, 256) // using small tmo buffer for demonstrating
	for {
		n, err := r.Read(tmp)
		if err != nil {
			return err
		}
		fmt.Printf("% x\n", tmp[:n])
	}
}
