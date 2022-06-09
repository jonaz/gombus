package gombus

import (
	"fmt"
	"io"
)

func RequestUD2(conn io.ReadWriter) error {
	data := []byte{
		0x10, // Start byte short frame
		0x7b, // REQ_UD2
		0x00, // device primary address
		0x00, // checksum
		0x16, // stop byte
	}
	_, err := conn.Write(data)
	return err
}

// water meter has 19004636 7704 14 07.
func SendUD2(conn io.ReadWriter) error {
	data := []byte{
		0x68, // Start byte long/control
		0x0b, // length
		0x0b, // length
		0x68, // Start byte long/control

		0x73, // REQ_UD2
		0xFD,
		0x52,

		0x00, // address
		0x00, // address
		0x00, // address
		0x00, // address

		0xFF, // manufacturer code
		0xFF, // manufacturer code

		0xFF, // id

		0xFF, // medium code

		0x00, // checksum
		0x16, // stop byte
	}

	data[15] = calcCheckSum(data[7:15])

	fmt.Printf("% x\n", data)
	_, err := conn.Write(data)
	return err
}

func calcCheckSum(data []byte) byte {
	var res byte
	for _, v := range data {
		res += v
	}
	return res & 0xff
}

// facit: 78 56 34 12 identification number = 12345678

// 1. Set the slave to primary address 8 without changing anything else:
// 68 06 06 68 | 53 FE 51 | 01 7A 08 | 25 16
