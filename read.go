package gombus

import (
	"bufio"
	"fmt"
	"io"
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
			if err != io.EOF {
				return err
			}
		}
		fmt.Printf("% x\n", tmp[:n])
	}
}
