// read2.go
package util_conn

import (
	"net"
	"time"
)

func read2(conn net.Conn, readTimeout int, w []byte) (err error) {
	err = conn.SetReadDeadline(time.Now().Add(time.Duration(readTimeout) * time.Millisecond))
	if err != nil {
		return err
	}
	_, err = conn.Read(w)
	if err != nil {
		return err
	}
	err = conn.SetReadDeadline(time.Time{})
	if err != nil {
		return err
	}
	return err
}
