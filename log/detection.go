// detection.go
package util_log

import (
	"io"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func OpenLogFile(path string, name string, writer string) *os.File {

	os.MkdirAll(path, os.ModePerm)
	hour := strconv.Itoa(time.Now().Hour())
	if time.Now().Hour() < 10 {
		hour = "0" + hour
	}
	minute := strconv.Itoa(time.Now().Minute())
	if time.Now().Minute() < 10 {
		minute = "0" + minute
	}
	second := strconv.Itoa(time.Now().Second())
	if time.Now().Second() < 10 {
		second = "0" + second
	}

	file, err := os.OpenFile(path+"/"+name+"."+strconv.Itoa(time.Now().Year())+"-"+strconv.Itoa(int(time.Now().Month()))+"-"+
		strconv.Itoa(time.Now().Day())+"_"+hour+minute+second+".log",
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

	}
	return file

}
func DetectionLogSize(file *os.File, size int64, path string, name string, writer string) {

	Ticker := time.NewTicker(time.Duration(1) * time.Minute) //单位分钟		1
	defer Ticker.Stop()

	for {
		select {
		case <-Ticker.C:
			fi, err := file.Stat()
			if err == nil {
				if fi.Size() > size*1024*1024 { //50M
					//关闭当前log文件，创建新log文件
					file.Close()
					file = OpenLogFile(path, name, writer)
				}
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
