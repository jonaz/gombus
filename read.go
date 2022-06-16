package gombus

import (
	"bufio"
	"fmt"
	"io"
)

func ReadLongFrame(r io.Reader) (LongFrame, error) {
	buf := bufio.NewReader(r)
	// TODO must check the length of the frame since 0x16 can occur anywhere :S
	// 68 LL LL 68
	// The L field gives the quantity of the user data inputs plus 3 (for C,A,CI)
	// AND add tests for this
	msg, err := buf.ReadBytes(0x16)
	if err != nil {
		return nil, err
	}
	return LongFrame(msg), nil
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
			if err != io.EOF {
				return err
			}
		}
		fmt.Printf("% x\n", tmp[:n])
	}
}
