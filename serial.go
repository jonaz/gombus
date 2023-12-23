package gombus

import serial "github.com/jonaz/serial"

func DialSerial(device string) (Conn, error) {
	// 2400, Even, 8, 1, None
	p, err := serial.OpenPort(&serial.Config{
		Name:     device,
		Baud:     2400,
		Size:     8,
		StopBits: 1,
		Parity:   serial.ParityEven,
	})

	return p, err
}
