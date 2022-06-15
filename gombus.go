package gombus

import (
	"encoding/binary"
	"fmt"
)

func RequestUD2(primaryID uint8) ShortFrame {
	data := NewShortFrame()
	data[2] = primaryID
	data.SetChecksum()
	return data
}

// water meter has 19004636 7704 14 07.
func SendUD2() Frame {
	data := []byte{
		0x68, // Start byte long/control
		0x00, // length
		0x00, // length
		0x68, // Start byte long/control

		0x73, // REQ_UD2
		0xFD,
		0x52, // CI-field selection of slave

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

	return data
}

func SetPrimaryUsingSecondary(secondary uint64, primary uint8) LongFrame {
	data := LongFrame{
		0x68, // Start byte long/control
		0x00, // length
		0x00, // length
		0x68, // Start byte long/control
		0x73, // REQ_UD2
		0xFD,
		0x51, // CI field data send
		0x00, // address
		0x00, // address
		0x00, // address
		0x00, // address
		0xFF, // manufacturer code
		0xFF, // manufacturer code
		0xFF, // id
		0xFF, // medium code
		0x01, // DIF field
		0x7a, // VIF field
		primary,
		0x00, // checksum
		0x16, // stop byte
	}

	a := UintToBCD(secondary, 4)
	data[7] = a[0]
	data[8] = a[1]
	data[9] = a[2]
	data[10] = a[3]

	data.Length()
	data.SetChecksum()
	return data
}

func SetPrimaryUsingPrimary(oldPrimary uint8, newPrimary uint8) LongFrame {
	data := LongFrame{
		0x68, // Start byte long/control
		0x06, // length
		0x06, // length
		0x68, // Start byte long/control
		0x73, // REQ_UD2
		oldPrimary,
		0x51, // CI field data send
		0x01, // DIF field
		0x7a, // VIF field
		newPrimary,
		0x00, // checksum
		0x16, // stop byte
	}

	data.Length()
	data.SetChecksum()
	return data
}

func calcCheckSum(data []byte) byte {
	var res byte
	for _, v := range data {
		res += v
	}
	return res & 0xff
}

func UintToBCD(value uint64, size int) []byte {
	buf := make([]byte, size)
	if value > 0 {
		remainder := value
		for pos := size - 1; pos >= 0 && remainder > 0; pos-- {
			tail := byte(remainder % 100)
			hi, lo := tail/10, tail%10
			buf[size-1-pos] = hi<<4 + lo
			remainder /= 100
		}
	}
	return buf
}

func BCDToInt(bcd []byte) int {
	var i = 0
	size := len(bcd)
	for k := range bcd {
		r0 := bcd[size-1-k] & 0xf
		r1 := bcd[size-1-k] >> 4 & 0xf
		r := r1*10 + r0
		i = i*100 + int(r)
	}
	return i
}

// if data fields is 0100 == 32 bit integer.
func Int32ToInt(data []byte) (int, error) {
	if len(data) != 4 {
		return 0.0, fmt.Errorf("wrong data length")
	}
	i := binary.LittleEndian.Uint32(data)
	return int(i), nil
}

func Int24ToInt(b []byte) int {
	_ = b[2] // bounds check hint to compiler; see golang.org/issue/14808
	return int(uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16)
}

func CheckKthBitSet(n, k int) bool {
	if n&(1<<(k)) == 0 {
		return false
	}
	return true
}

/*

func AsRoundedFloat(data []byte) (float64, error) {
	// this works OK according to protocol example:
	// Extract temperature 21,8
	// 0000   3d 05 00 d3 5e ae 41 67 3e
	if len(data) != 4 {
		return 0.0, fmt.Errorf("could not pase float from binary")
	}
	bits := binary.LittleEndian.Uint32(data)
	float := float64(math.Float32frombits(bits))
	return math.Round(float*100) / 100, nil
}
*/
