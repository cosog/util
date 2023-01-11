// detection.go
package util_log

import (
	"io"
	"io/fs"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func OpenLogFile(path string, name string, writer string) (*os.File, fs.FileInfo) {

	var fi fs.FileInfo
	os.MkdirAll(path, os.ModePerm)
	file, err := os.OpenFile(path+name+"_"+strconv.Itoa(time.Now().Year())+strconv.Itoa(int(time.Now().Month()))+strconv.Itoa(time.Now().Day())+
		strconv.Itoa(time.Now().Hour())+strconv.Itoa(time.Now().Minute())+strconv.Itoa(time.Now().Second())+".log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file:", err)

	} else {
		switch strings.ToLower(writer) {
		case "file":
			log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
			log.SetOutput(file)
		case "stdout":
			log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
			log.SetOutput(os.Stderr)
		case "stdout+file":
			fallthrough
		case "file+stdout":
			log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
			log.SetOutput(io.MultiWriter(file, os.Stderr))
		case "close":
			log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
			log.SetOutput(io.MultiWriter())
		}

		fi, _ = file.Stat()
	}
	return file, fi

}
func DetectionLogSize(file *os.File, fi fs.FileInfo, path string, name string, writer string) {
	Ticker := time.NewTicker(time.Duration(1) * time.Minute) //单位分钟		1
	defer Ticker.Stop()

	for {
		select {
		case <-Ticker.C:
			if fi.Size() > 50*1024*1024 { //50M
				//关闭当前log文件，创建新log文件
				file.Close()
				file, fi = OpenLogFile(path, name, writer)
			}
		}
		runtime.Gosched()
	}
}
func DetectionLogModTime(path string, name string) {
	Ticker := time.NewTicker(time.Duration(24) * time.Hour) //小时
	defer Ticker.Stop()

	for {
		select {
		case <-Ticker.C:

		}
		runtime.Gosched()
	}
}
