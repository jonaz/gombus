package gombus

import (
	"context"
	"fmt"
	"net"
	"time"
)

func Dial(addr string) (net.Conn, error) {
	var dialer net.Dialer
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	conn, err := dialer.DialContext(ctx, "tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to dial: %w", err)
	}
	return conn, err
}
