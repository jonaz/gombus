package gombus

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
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

func (sf ShortFrame) SetChecksum() {
	size := len(sf)
	sf[size-2] = calcCheckSum(sf[1 : size-2])
}

func (sf ShortFrame) SetAddress(primary uint8) {
	sf[2] = primary
}
func (sf ShortFrame) SetC(c uint8) {
	sf[1] = c
}

func (sf ShortFrame) C() C {
	return C(sf[1])
}

func (sf ShortFrame) SetFCB() {
	sf[1] |= CONTROL_MASK_FCB
}

func (sf ShortFrame) SetFCV() {
	sf[1] |= CONTROL_MASK_FCV
}

func (sf ShortFrame) ClearFCB() {
	sf[1] &^= CONTROL_MASK_FCB
}

func (sf ShortFrame) ClearFCV() {
	sf[1] &^= CONTROL_MASK_FCV
}

func (sf ShortFrame) A() byte {
	return sf[2]
}

type LongFrame Frame

func (lf LongFrame) SetChecksum() {
	size := len(lf)
	lf[size-2] = calcCheckSum(lf[4 : size-2])
}

func (lf LongFrame) SetLength() {
	lf[1] = byte(len(lf) - 6)
	lf[2] = byte(len(lf) - 6)
}

func (lf LongFrame) L() int {
	return int(lf[1])
}

func (lf LongFrame) C() C {
	return C(lf[4])
}

func (lf LongFrame) SetFCB() {
	lf[4] |= CONTROL_MASK_FCB
}

func (lf LongFrame) SetFCV() {
	lf[4] |= CONTROL_MASK_FCV
}

func (lf LongFrame) ClearFCB() {
	lf[4] &^= CONTROL_MASK_FCB
}

func (lf LongFrame) ClearFCV() {
	lf[4] &^= CONTROL_MASK_FCV
}

func (lf LongFrame) A() byte {
	return lf[5]
}
func (lf LongFrame) CI() byte {
	return lf[6]
}

func (lf LongFrame) Decode() (*DecodedFrame, error) {
	if lf.CI() != 0x72 {
		return nil, fmt.Errorf("unknown longframe, only supports variable data response for now")
	}

	man, err := lf.DecodeManufacturer()
	if err != nil {
		return nil, err
	}

	dr, err := lf.decodeData(lf[19 : len(lf)-2])
	if err != nil {
		return nil, err
	}
	dt, err := DeviceTypeLookup(lf[14])
	if err != nil {
		return nil, err
	}

	dFrame := &DecodedFrame{
		SerialNumber: BCDToInt(lf[7:11]),
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

func (lf LongFrame) decodeData(data []byte) ([]DecodedDataRecord, error) {
	records := make([]DecodedDataRecord, 0)

	var dData DecodedDataRecord
	// var dif byte
	dif := -1
	var dife []byte
	lookForData := false
	lookForDIFE := false
	lookForVIF := false
	lookForVIFE := false
	remainingData := 0
	customUnit := ""
	var vife []byte
	var vif byte
	for i, v := range data {
		if remainingData > 0 {
			remainingData--
			continue
		}
		// expect first one is a DIF
		if dif == -1 {
			// TODO handle special functionsFCB-
			// DIF	Function
			// 0Fh	Start of manufacturer specific data structures to end of user data
			// 1Fh DONE	Same meaning as DIF = 0Fh + More records follow in next telegram
			// 2Fh	Idle Filler (not to be interpreted), following byte = DIF
			// 3Fh..6Fh	Reserved
			// 7Fh	Global readout request (all storage#, units, tariffs, function fields)

			dData = DecodedDataRecord{}
			dData.Function = DecodeRecordFunction(v)
			dData.StorageNumber = int(v) & DATA_RECORD_DIF_MASK_STORAGE_NO

			// 1Fh DONE Same meaning as DIF = 0Fh + More records follow in next telegram
			if v == 0x1f {
				dData.HasMoreRecords = true
				records = append(records, dData)
			}
			// 2Fh  Idle Filler (not to be interpreted), following byte = DIF
			if v == 0x2f {
				continue
			}

			dif = int(v)
			logrus.Debugf("dif is: % x\n", dif)
			if CheckKthBitSet(int(v), 7) {
				lookForDIFE = true
				continue
			}
			lookForVIF = true
			continue
		}
		if lookForDIFE { // has another DIFE{
			dife = append(dife, v)
			// TODO validate we dont have more than 10 here
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

			// In case of VIF = 7Ch / FCh the true VIF is represented by the following ASCII string with the length given in the first byte
			if vif == 0x7c {
				l := int(data[i+1])
				customUnit = DecodeASCII(data[i+2 : i+2+l])
				remainingData = l + 1
				lookForVIF = false
				lookForData = true
				continue
			}

			lookForVIF = false
			lookForData = true
			continue
		}
		if lookForVIFE {
			vife = append(vife, v)
			// TODO validate we dont have more than 10 here
			if CheckKthBitSet(int(v), 7) {
				// lookForVIFE = true
				continue
			}
			lookForVIFE = false
			lookForData = true
			continue
		}

		if lookForData {
			if customUnit != "" {
				dData.Unit = VIF{
					Exp:         1,
					Unit:        customUnit,
					Type:        VIFUnit["VARIABLE_VIF"],
					VIFUnitDesc: "",
				}
				customUnit = ""
			} else {
				dData.Unit = DecodeUnit(vif, vife)
			}
			dData.StorageNumber = DecodeStorageNumber(dif, dife)
			dData.Device = DecodeDevice(dif, dife)
			dData.Tariff = DecodeTariff(dif, dife)

			difCoding := dif & DATA_RECORD_DIF_MASK_DATA
			logrus.Debugf("Datarecord dif mask: %b ( %#x )", difCoding, difCoding)

			switch difCoding {
			// 0000	No data
			case 0x00:
				remainingData = 0

			// 0001	8 Bit Integer
			case 0x01:
				remainingData = 0
				dData.RawValue = float64(data[i])
				logrus.Debugf("data dif 0x01 is: % x\n", data[i])

			// 0010	16 Bit Integer
			case 0x02:
				remainingData = 1
				dData.RawValue = float64(binary.LittleEndian.Uint16(data[i : i+2]))
				logrus.Debugf("data dif 0x02 is: % x\n", data[i:i+4])

			// 0011	24 Bit Integer
			case 0x03:
				remainingData = 2

			// 4 byte (32 bit)
			case 0x04:
				remainingData = 3
				logrus.Debugf("data dif 0x04 is: % x\n", data[i:i+4])
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
				// With data field = 1101b several data types with variable length can be used. The length of the data is given with the first byte of data, which is here called LVAR.
				// LVAR = 00h .. BFh : ASCII string with LVAR characters
				// LVAR = C0h .. CFh : positive BCD number with (LVAR - C0h) · 2 digits
				// LVAR = D0h .. DFH : negative BCD number with (LVAR - D0h) · 2 digits
				// LVAR = E0h .. EFh : binary number with (LVAR - E0h) bytes
				// LVAR = F0h .. FAh : floating point number with (LVAR - F0h) bytes [to be defined]
				// LVAR = FBh .. FFh : Reserved
				size := 0
				if data[i] <= 0xBF {
					size = int(data[i])
					dData.ValueString = DecodeASCII(data[i+1 : i+1+size])
				} else if data[i] >= 0xC0 && data[i] <= 0xCF {
					size = (int(data[i]) - 0xC0) * 2
					// TODO data here
				} else if data[i] >= 0xD0 && data[i] <= 0xDF {
					size = (int(data[i]) - 0xD0) * 2
					// TODO data here
				} else if data[i] >= 0xE0 && data[i] <= 0xEF {
					size = int(data[i]) - 0xE0
					// TODO data here
				} else if data[i] >= 0xF0 && data[i] <= 0xFA {
					size = int(data[i]) - 0xF0
					// TODO data here
				}
				remainingData = size

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
			dife = nil
			vif = 0
			vife = nil
			logrus.Debug("rawValue: ", dData.RawValue)
			logrus.Debug("valueString: ", dData.ValueString)
			dData.Value = dData.Unit.Value(dData.RawValue)
			records = append(records, dData)
		}
	}

	return records, nil
}

func (lf LongFrame) DecodeManufacturer() (string, error) {
	id := int(binary.LittleEndian.Uint16(lf[11:13]))
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

	Unit     VIF
	Exponent float64
	Type     string
	Quantity string

	Value       float64
	ValueString string
	RawValue    float64

	HasMoreRecords bool
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

func (df DecodedFrame) HasMoreRecords() bool {
	if len(df.DataRecords) == 0 {
		return false
	}
	return df.DataRecords[len(df.DataRecords)-1].HasMoreRecords
}

const SingleCharacterFrame = 0xe5

type C byte

// FCB Frame Count-Bit.
func (c C) FCB() bool {
	return (c & CONTROL_MASK_FCB) > 0
}

// FCV Frame Count Valid indicates we want to use frame counting in the following request/responses.
func (c C) FCV() bool {
	return (c & CONTROL_MASK_FCV) > 0
}
