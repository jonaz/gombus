package gombus

import (
	"encoding/hex"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToBCD(t *testing.T) {
	// facit: 78 56 34 12 identification number = 12345678
	s := fmt.Sprintf("% x", UintToBCD(12345678, 4))
	assert.Equal(t, "78 56 34 12", s)
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
func TestDecodeLongFrame(t *testing.T) {
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
		82 c0 40 fd ba 73 00 00 1f ef 16`
	// 0x84 == 10000100 , 0100 == 32 bit integer // https://m-bus.com/documentation-wired/06-application-layer
	// 0x40 == 01000000 ..TODO ..TODO
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "\n", "")
	data, err := hex.DecodeString(s)
	assert.NoError(t, err)
	fmt.Println(data)
}
