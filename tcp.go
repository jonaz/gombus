package gombus

import (
	"context"
	"fmt"
	"net"
	"time"
)

type Conn struct {
	conn net.Conn
}

func Dial(addr string) (*Conn, error) {
	var dialer net.Dialer
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	conn, err := dialer.DialContext(ctx, "tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to dial: %w", err)
	}
	return &Conn{conn: conn}, nil
}

func (c *Conn) Read(b []byte) (n int, err error) {
	return c.conn.Read(b)
}
func (c *Conn) Write(b []byte) (n int, err error) {
	return c.conn.Write(b)
}
func (c *Conn) Close() error {
	return c.conn.Close()
}
func (c *Conn) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}
func (c *Conn) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}
func (c *Conn) SetDeadline(t time.Time) error {
	return c.conn.SetDeadline(t)
}
func (c *Conn) SetReadDeadline(t time.Time) error {
	return c.conn.SetReadDeadline(t)
}
func (c *Conn) SetWriteDeadline(t time.Time) error {
	return c.conn.SetWriteDeadline(t)
}

var ErrNoLongFrameFound = fmt.Errorf("no long frame found")

func (c *Conn) ReadLongFrame() (LongFrame, error) {
	buf := make([]byte, 4096)
	tmp := make([]byte, 4096)

	// foundStart := false
	length := 0
	globalN := -1
	for {
		err := c.conn.SetReadDeadline(time.Now().Add(2 * time.Second))
		if err != nil {
			return LongFrame{}, fmt.Errorf("error from SetReadDeadline: %w", err)
		}

		n, err := c.conn.Read(tmp)
		if err != nil {
			return LongFrame{}, fmt.Errorf("error reading from tcp connection: %w", err)
		}

		for _, b := range tmp[:n] {
			globalN++
			buf[globalN] = b

			if globalN > 256 {
				return LongFrame{}, ErrNoLongFrameFound
			}

			// look for end byte after length +C+A+CI+checksum
			if length != 0 && globalN > length+4 && b == 0x16 {
				return LongFrame(buf[:globalN+1]), nil
			}

			// look for start sequence 68 LL LL 68
			if length == 0 && buf[0] == 0x68 && buf[3] == 0x68 && buf[1] == buf[2] {
				length = int(buf[1])
			}
		}
	}
}
