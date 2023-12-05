package gombus

import "time"

type Conn interface {
	Read(b []byte) (n int, err error)
	Write(b []byte) (n int, err error)
	SetReadDeadline(t time.Time) error
	Close() error
}
