// write2
package util_conn

import (
	"net"
	"time"
)

func Write2(conn net.Conn, writeTimeout int, w []byte) (n int, err error) {
	err = conn.SetWriteDeadline(time.Now().Add(time.Duration(writeTimeout) * time.Millisecond))
	if err != nil {
		return 0, err
	}
	n, err = conn.Write(w)
	if err != nil {
		return n, err
	}
	err = conn.SetWriteDeadline(time.Time{})
	if err != nil {
		return 0, err
	}
	return n, err
}
