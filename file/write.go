// write
package util_file

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
)

func WriteFile(str string, w interface{}) {

	f, err := os.OpenFile(str, os.O_RDWR|os.O_CREATE|os.O_SYNC, 0666)
	if err != nil {
		fmt.Println("open file fail:", err)
		return
	}
	defer f.Close()
	f.Truncate(0)

	a, err := json.MarshalIndent(w, "", "	")

	start := 0
	for {
		n, _ := f.Write(a[start:])
		if n == len(a)-start {
			break
		} else {
			start += n
		}
		runtime.Gosched()
	}

	// f.Close()
}
func WriteFileZip(str string, w interface{}) {

	f, err := os.OpenFile(str, os.O_RDWR|os.O_CREATE|os.O_SYNC, 0666)
	if err != nil {
		fmt.Println("open file fail:", err)
		return
	}
	defer f.Close()
	f.Truncate(0)

	// a, err := json.MarshalIndent(w, "", "	")
	a, err := json.Marshal(w)

	start := 0
	for {
		n, _ := f.Write(a[start:])
		if n == len(a)-start {
			break
		} else {
			start += n
		}
		runtime.Gosched()
	}

	// f.Close()
}
