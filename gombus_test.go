package gombus

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	logrus.SetLevel(logrus.DebugLevel)
	os.Exit(m.Run())
}

func TestToBCD(t *testing.T) {
	// facit: 78 56 34 12 identification number = 12345678
	s := fmt.Sprintf("% x", UintToBCD(12345678, 4))
	assert.Equal(t, "78 56 34 12", s)
}
func TestFromBCD(t *testing.T) {
	// facit: 78 56 34 12 identification number = 12345678
	h, err := hex.DecodeString("78563412")
	assert.NoError(t, err)

	i := BCDToInt(h)
	assert.Equal(t, 12345678, i)
}

func TestCheckKthBitSet(t *testing.T) {
	assert.True(t, CheckKthBitSet(0x80, 7))
	assert.False(t, CheckKthBitSet(0xf, 7))
	assert.False(t, CheckKthBitSet(0x2a, 7))
	assert.True(t, CheckKthBitSet(0x40, 6))
}

func TestPrimaryUsingSecondary(t *testing.T) {
	data := SetPrimaryUsingSecondary(19004636, 2)
	s := fmt.Sprintf("% x", data)
	assert.Equal(t, "68 0e 0e 68 73 fd 51 36 46 00 19 ff ff ff ff 01 7a 02 cf 16", s)
}

func TestPrimaryUsingPrimary(t *testing.T) {
	data := SetPrimaryUsingPrimary(0, 3)
	s := fmt.Sprintf("% x", data)
	assert.Equal(t, "68 06 06 68 73 00 51 01 7a 03 42 16", s)
}
func TestDecodeLongFrameGAROSecondFrame(t *testing.T) {
	// Response from garo electric meter
	s := `
		68 78 78 68 
		08 01 72 
		14 21 07 90
		36 1c 
		c7 
		02 
		25 
		00 
		00 00 
		84 40 2a a0 09 00 00
		84 80 40 2a ba 00 00 00 
		84 c0 40 2a 00 00 00 00 
		84 40 fb 97 72 fb fe ff ff 
		84 80 40 fb 97 72 4b 00 00 00 
		84 c0 40 fb 97 72 00 00 00 00 
		84 40 fb b7 72 ae 09 00 00 
		84 80 40 fb b7 72 c8 00 00 00 
		84 c0 40 fb b7 72 00 00 00 00 
		82 40 fd ba 73 e2 03 
		82 80 40 fd ba 73 9f 03 
		82 c0 40 fd ba 73 00 00 1f 
		ef 16`
	// 0x84 == 10000100 , 1000 = extension bit 1,  0100 == 32 bit integer // https://m-bus.com/documentation-wired/06-application-layer
	// DIF:
	// 7 extension bit
	// 6 LSB always 0
	// 5 function field
	// 4 function field (2 bytes) 00 == Instantaneous Value
	// 3 data field for example 0100 == 32 bit integer
	// 2 data field
	// 1 data field
	// 0 data field

	// DIFE
	// 0x40 == 01000000
	// 0x80 == 10000000
	// 0xc0 == 11000000
	// 7 extension bit
	// 6 (device) unit 0 = reactive, 1=apparent
	// 5 tariff
	// 4 tariff
	// 3 storagenumber
	// 2 storagenumber
	// 1 storagenumber
	// 0 storagenumber

	// VIF
	// 0x2a == 00101010 == E010 1010	Reserved
	// 7 extension bit
	// 6 unit and multiplier
	// 5 unit and multiplier
	// 4 unit and multiplier
	// 3 unit and multiplier
	// 2 unit and multiplier
	// 1 unit and multiplier
	// 0 unit and multiplier 7 bits total

	// VIFE
	// inget för extension bit == 0
	// DATA
	// a0 09 00 00

	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.ReplaceAll(s, "\t", "")
	data, err := hex.DecodeString(s)
	assert.NoError(t, err)

	frame := LongFrame(data)

	fmt.Println(frame)
	dFrame, err := frame.Decode()
	assert.NoError(t, err)
	assert.Equal(t, 90072114, dFrame.SerialNumber)

	// fmt.Printf("%#v\n", dFrame)
	// spew.Dump(dFrame)
}
func TestDecodeLongFrameGAROFirstFrame(t *testing.T) {
	s := `68 65 65 68 08 01 72 14 21 07 90 36 1c c7 02 4d 00 00 00 04 05 9c 31 01 00 04 fb 82 75 63 91 00 00 04 2a 36 08 00 00 04 fb 97 72 ca fe ff ff 04 fb b7 72 6d 08 00 00 02 fd ba 73 dc 03 84 80 80 40 fd 48 c4 0f 00 00 04 fd 48 1a 09 00 00 84 40 fd 59 d2 04 00 00 84 80 40 fd 59 78 00 00 00 84 c0 40 fd 59 00 00 00 00 1f 95 16`
	s = strings.ReplaceAll(s, " ", "")
	data, err := hex.DecodeString(s)
	assert.NoError(t, err)

	frame := LongFrame(data)

	fmt.Println(frame)
	dFrame, err := frame.Decode()
	assert.NoError(t, err)
	assert.Equal(t, 90072114, dFrame.SerialNumber)

	// fmt.Printf("%#v\n", dFrame)
	// spew.Dump(dFrame)
}

// TODO test this frame which has 0x16 in the middle. We must read the whole one.
// 68 56 56 68 08 02 72 36 46 00 19 77 04 14 07 40 10 00 00 0C 78 36 46 00 19 0D 7C 08 44 49 20 2E 74 73 75 63 0A 20 20 20 20 20 20 20 20 20 20 04 6D 32 16 D0 26 02 7C 09 65 6D 69 74 20 2E 74 61 62 9A 10 04 13 75 68 03 00 04 93 7F 00 00 00 00 44 13 27 51 03 00 0F 00 00 1F A6 16

func TestInt24ToInt(t *testing.T) {
	// 03 13 15 31 00 Data block 1: unit 0, storage No 0, no tariff, instantaneous volume, 12565 l (24 bit integer)
	d := []byte{0x15, 0x31, 0x0}
	res := Int24ToInt(d)
	assert.Equal(t, 12565, res)
}

// Example for a RSP_UD with variable data structure answer (mode 1):
// (all values are hex.)

// 68 1F 1F 68 header of RSP_UD telegram (length 1Fh=31d bytes)
// 08 02 72 C field = 08 (RSP), address 2, CI field 72H (var.,LSByte first)
// 78 56 34 12 identification number = 12345678
// 24 40 01 07 manufacturer ID = 4024h (PAD in EN 61107), generation 1, water
// 55 00 00 00 TC = 55h = 85d, Status = 00h, Signature = 0000h
// 03 13 15 31 00 Data block 1: unit 0, storage No 0, no tariff, instantaneous volume, 12565 l (24 bit integer)
// DA 02 3B 13 01 Data block 2: unit 0, storage No 5, no tariff, maximum volume flow, 113 l/h (4 digit BCD)
// 8B 60 04 37 18 02 Data block 3: unit 1, storage No 0, tariff 2, instantaneous energy, 218,37 kWh (6 digit BCD)
// 18 16 checksum and stopsign
