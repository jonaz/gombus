package gombus

import (
	"encoding/hex"
	"net"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadLongFrame(t *testing.T) {
	s := "68 56 56 68 08 02 72 36 46 00 19 77 04 14 07 40 10 00 00 0c 78 36 46 00 19 0d 7c 08 44 49 20 2e 74 73 75 63 0a 20 20 20 20 20 20 20 20 20 20 04 6d 32 16 d0 26 02 7c 09 65 6d 69 74 20 2e 74 61 62 9a 10 04 13 75 68 03 00 04 93 7f 00 00 00 00 44 13 27 51 03 00 0f 00 00 1f a6 16 00 00 00"
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "", "")
	data, err := hex.DecodeString(s)
	assert.NoError(t, err)

	server, client := net.Pipe()
	conn := &conn{conn: client}
	go func() {
		// Do some stuff
		for _, d := range data {
			_, err := server.Write([]byte{d})
			assert.NoError(t, err)
		}

		server.Close()
	}()

	frame, err := ReadLongFrame(conn)
	assert.NoError(t, err)

	// make sure we read the whole frame correctly!
	assert.Len(t, frame, 92)

	dFrame, err := frame.Decode()
	assert.NoError(t, err)

	// water 1243 m2!
	assert.Equal(t, 217.383, dFrame.DataRecords[6].Value)

	// spew.Dump(dFrame)
}
