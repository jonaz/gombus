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
	for i, v := range data {
		// expect first one is a DIF
		if dif == -1 {
			dData = DecodedDataRecord{}
			dif = int(v)
			dData.Function = DecodeRecordFunction(v)

			if CheckKthBitSet(int(v), 7) {
				lookForDIFE = true
			}

			lookForVIF = true

			continue
		}
		if lookForDIFE { // has another DIFE{
			// dife = int(v)
			// TODO read device, tariff and StorageNumber
			if CheckKthBitSet(int(v), 7) {
				lookForDIFE = true
				continue
			}
			lookForDIFE = false
			lookForVIF = true
			continue
		}

		// E111 1100 7 bits data
		if lookForVIF {
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
			if CheckKthBitSet(int(v), 7) {
				lookForVIFE = true
				continue
			}
			lookForVIFE = false
			lookForData = true
			continue
		}

		if lookForData {
			switch dif & DATA_RECORD_DIF_MASK_DATA {
			// 4 byte (32 bit)
			case 0x04:
				fmt.Printf("0x04 is: % x\n", data[i:i+4])
				v, err := Int32ToInt(data[i : i+4])
				if err != nil {
					return nil, err
				}

				dData.RawValue = float64(v)
			}
			lookForData = false
			dif = -1
			records = append(records, dData)
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
