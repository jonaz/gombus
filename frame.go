package gombus

import (
	"encoding/binary"
	"fmt"
	"time"
)

type Frame []byte

type ShortFrame Frame

func NewShortFrame() ShortFrame {
	return ShortFrame{
		0x10, // Start byte short frame
		0x7b, // C field
		0x00, // A field
		0x00, // checksum
		0x16, // stop byte
	}
}

func (cf ShortFrame) SetChecksum() {
	size := len(cf)
	cf[size-2] = calcCheckSum(cf[1 : size-2])
}

type LongFrame Frame

func (cf LongFrame) SetChecksum() {
	size := len(cf)
	cf[size-2] = calcCheckSum(cf[4 : size-2])
}

func (cf LongFrame) Length() {
	cf[1] = byte(len(cf) - 6)
	cf[2] = byte(len(cf) - 6)
}

func (cf LongFrame) L() int {
	return int(cf[1])
}

func (cf LongFrame) C() byte {
	return cf[4]
}

func (cf LongFrame) A() byte {
	return cf[5]
}
func (cf LongFrame) CI() byte {
	return cf[6]
}

func (cf LongFrame) Decode() (*DecodedFrame, error) {
	if cf.CI() != 0x72 {
		return nil, fmt.Errorf("unknown longframe, only supports variable data response for now")
	}

	man, err := cf.DecodeManufacturer()
	if err != nil {
		return nil, err
	}

	dr, err := cf.decodeData(cf[19 : len(cf)-2])
	if err != nil {
		return nil, err
	}
	dt, err := DeviceTypeLookup(cf[14])
	if err != nil {
		return nil, err
	}

	dFrame := &DecodedFrame{
		SerialNumber: BCDToInt(cf[7:11]),
		Manufacturer: man,
		ProductName:  "", // TODO
		Version:      0,  // TODO
		DeviceType:   dt,
		AccessNumber: 0, // TODO
		Signature:    0, // TODO
		Status:       0, // TODO
		DataRecords:  dr,
	}

	return dFrame, nil
}

func (cf LongFrame) decodeData(data []byte) ([]DecodedDataRecord, error) {
	records := make([]DecodedDataRecord, 0)

	var dData DecodedDataRecord
	// var dif byte
	dif := -1
	lookForData := false
	lookForDIFE := false
	lookForVIF := false
	lookForVIFE := false
	remainingData := 0
	var vife []byte
	var vif byte
	for i, v := range data {
		if remainingData > 0 {
			remainingData--
			continue
		}
		// expect first one is a DIF
		if dif == -1 {
			dData = DecodedDataRecord{}
			dif = int(v)
			dData.Function = DecodeRecordFunction(v)
			fmt.Printf("dif is: % x\n", dif)
			if CheckKthBitSet(int(v), 7) {
				lookForDIFE = true
				continue
			}
			lookForVIF = true
			continue
		}
		if lookForDIFE { // has another DIFE{
			// dife = int(v)
			// TODO read device, tariff and StorageNumber
			if CheckKthBitSet(int(v), 7) {
				// lookForDIFE = true
				continue
			}
			lookForDIFE = false
			lookForVIF = true
			continue
		}

		// E111 1100 7 bits data
		if lookForVIF {
			vif = v
			if CheckKthBitSet(int(v), 7) {
				lookForVIF = false
				lookForVIFE = true
				continue
			}
			lookForVIF = false
			lookForData = true
			continue
		}
		if lookForVIFE {
			vife = append(vife, v)
			if CheckKthBitSet(int(v), 7) {
				// lookForVIFE = true
				continue
			}
			lookForVIFE = false
			lookForData = true
			continue
		}

		if lookForData {
			fmt.Printf("unit: %#v\n", DecodeUnit(vif, vife))
			switch dif & DATA_RECORD_DIF_MASK_DATA {
			// 0000	No data
			case 0x00:
				remainingData = 0

			// 0001	8 Bit Integer
			case 0x01:
				remainingData = 0
				dData.RawValue = float64(data[i])
				fmt.Printf("data dif 0x01 is: % x\n", data[i])

			// 0010	16 Bit Integer
			case 0x02:
				remainingData = 1
				dData.RawValue = float64(binary.LittleEndian.Uint16(cf[11:13]))
				fmt.Printf("data dif 0x02 is: % x\n", data[i:i+4])

			// 0011	24 Bit Integer
			case 0x03:
				remainingData = 2

			// 4 byte (32 bit)
			case 0x04:
				remainingData = 3
				fmt.Printf("data dif 0x04 is: % x\n", data[i:i+4])
				v, err := Int32ToInt(data[i : i+4])
				if err != nil {
					return nil, err
				}

				dData.RawValue = float64(v)

			// 0101	32 Bit Real
			case 0x05:
				remainingData = 3

			// 0110	48 Bit Integer
			case 0x06:
				remainingData = 5

			// 0111	64 Bit Integer
			case 0x07:
				remainingData = 7

			// 1000	Selection for Readout
			case 0x08:
				remainingData = 0

			// 1001	2 digit BCD
			case 0x09:
				remainingData = 0
				dData.RawValue = float64(BCDToInt(data[i : i+1]))

			// 1010	4 digit BCD
			case 0x0a:
				remainingData = 1
				dData.RawValue = float64(BCDToInt(data[i : i+2]))

			// 1011	6 digit BCD
			case 0x0b:
				remainingData = 2
				dData.RawValue = float64(BCDToInt(data[i : i+3]))

			// 1100	8 digit BCD
			case 0x0c:
				remainingData = 3
				dData.RawValue = float64(BCDToInt(data[i : i+4]))

			// 1101	variable length
			case 0x0d:
				remainingData = 0 // TODO what here?

			// 1110	12 digit BCD
			case 0x0e:
				remainingData = 5
				dData.RawValue = float64(BCDToInt(data[i : i+6]))

			// 1111	Special Functions
			case 0x0f:
				remainingData = 0 // TODO what here?
			}
			lookForData = false
			dif = -1
			vif = 0
			vife = nil
			fmt.Println("rawValue", dData.RawValue)
			records = append(records, dData)
			fmt.Println("")
		}
	}

	return records, nil
}

func (cf LongFrame) DecodeManufacturer() (string, error) {
	id := int(binary.LittleEndian.Uint16(cf[11:13]))
	return fmt.Sprintf(
		"%c%c%c",
		rune(((id>>10)&0x001F)+64),
		rune(((id>>5)&0x001F)+64),
		rune((id&0x001F)+64),
	), nil
}

type DecodedDataRecord struct {
	Function      string
	StorageNumber int

	Tariff int
	Device int

	Unit     string
	Exponent float64
	Type     string
	Quantity string

	Value    string
	RawValue float64
}

type DecodedFrame struct {
	SerialNumber int
	Manufacturer string
	ProductName  string
	Version      int
	DeviceType   string
	AccessNumber int16

	Signature int16

	Status int
	// ReadableStatus string // TODO make function on struct!

	DataRecords []DecodedDataRecord

	ParsedAt time.Time
}

const SingleCharacterFrame = 0xe5
