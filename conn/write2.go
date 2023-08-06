// write2
package util_conn

import (
	"net"
	"time"
)

func Write2(conn net.Conn, writeTimeout int, w []byte) (err error) {
	err = conn.SetWriteDeadline(time.Now().Add(time.Duration(writeTimeout) * time.Millisecond))
	if err != nil {
		return err
	}
	_, err = conn.Write(w)
	if err != nil {
		return err
	}
	err = conn.SetWriteDeadline(time.Time{})
	if err != nil {
		return err
	}
	return err
}
