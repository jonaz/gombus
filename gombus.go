package gombus

type Frame []byte

type ShortFrame Frame

func (cf ShortFrame) Checksum() {
	size := len(cf)
	cf[size-2] = calcCheckSum(cf[1 : size-2])
}

type ControlFrame Frame

func (cf ControlFrame) Checksum() {
	size := len(cf)
	cf[size-2] = calcCheckSum(cf[4 : size-2])
}

func (cf ControlFrame) Length() {
	cf[1] = byte(len(cf) - 6)
	cf[2] = byte(len(cf) - 6)
}

type LongFrame ControlFrame

func NewShortFrame() ShortFrame {
	return ShortFrame{
		0x10, // Start byte short frame
		0x7b, // C field
		0x00, // A field
		0x00, // checksum
		0x16, // stop byte
	}
}

const SingleCharacterFrame = 0xe5

func RequestUD2(primaryID uint8) ShortFrame {
	data := NewShortFrame()
	data[2] = primaryID
	data.Checksum()
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

func SetPrimaryUsingSecondary(secondary uint64, primary uint8) ControlFrame {
	data := ControlFrame{
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
	data.Checksum()
	return data
}

func SetPrimaryUsingPrimary(oldPrimary uint8, newPrimary uint8) ControlFrame {
	data := ControlFrame{
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
	data.Checksum()
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
