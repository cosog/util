// read2.go
package util_conn

import (
	"net"
	"time"
)

func Read2(conn net.Conn, readTimeout int, r *[]byte) (n int, err error) {
	err = conn.SetReadDeadline(time.Now().Add(time.Duration(readTimeout) * time.Millisecond))
	if err != nil {
		return 0, err
	}
	n, err = conn.Read(*r)
	if err != nil {
		return n, err
	}
	err = conn.SetReadDeadline(time.Time{})
	if err != nil {
		return n, err
	}
	return n, err
}
