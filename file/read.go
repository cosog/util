// read
package util_file

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"runtime"
)

func ReadFile(str string, r interface{}) {

	f, err := os.OpenFile(str, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println("open file fail:", err)

	}
	defer f.Close()

	f.Seek(0, os.SEEK_SET)

	var s []byte
	for {
		r := make([]byte, 128)

		n, err := f.Read(r)
		if err != nil && err != io.EOF {
			fmt.Println("err:", err)
		}
		if n != 0 {
			s = append(s, r[0:n]...)
		} else {
			break
		}
		runtime.Gosched()
	}

	json.Unmarshal(s, r)

	f.Close()
}
