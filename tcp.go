package gombus

import (
	"context"
	"fmt"
	"net"
	"time"
)

type conn struct {
	conn net.Conn
}

func Dial(addr string) (Conn, error) {
	var dialer net.Dialer
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	c, err := dialer.DialContext(ctx, "tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to dial: %w", err)
	}
	return &conn{conn: c}, nil
}

func (c *conn) Read(b []byte) (n int, err error) {
	return c.conn.Read(b)
}
func (c *conn) Write(b []byte) (n int, err error) {
	return c.conn.Write(b)
}
func (c *conn) Close() error {
	return c.conn.Close()
}
func (c *conn) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}
func (c *conn) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}
func (c *conn) SetDeadline(t time.Time) error {
	return c.conn.SetDeadline(t)
}
func (c *conn) SetReadDeadline(t time.Time) error {
	return c.conn.SetReadDeadline(t)
}
func (c *conn) SetWriteDeadline(t time.Time) error {
	return c.conn.SetWriteDeadline(t)
}

func (c *conn) Conn() net.Conn {
	return c.conn
}
