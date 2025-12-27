// detection.go
package util_log

import (
	"io"
	"io/ioutil"
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

	switch strings.ToLower(writer) {
	case "file":
		file, err := os.OpenFile(path+"/"+name+"."+strconv.Itoa(time.Now().Year())+"-"+strconv.Itoa(int(time.Now().Month()))+"-"+
			strconv.Itoa(time.Now().Day())+"_"+hour+minute+second+".log",
			os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Println("Failed to open log file:", err)
			log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
			log.SetOutput(io.MultiWriter())
		} else {
			log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
			log.SetOutput(file)
		}
		return file
	case "stdout":
		log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
		log.SetOutput(os.Stderr)
		return nil
	case "stdout+file":
		fallthrough
	case "file+stdout":
		file, err := os.OpenFile(path+"/"+name+"."+strconv.Itoa(time.Now().Year())+"-"+strconv.Itoa(int(time.Now().Month()))+"-"+
			strconv.Itoa(time.Now().Day())+"_"+hour+minute+second+".log",
			os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Println("Failed to open log file:", err)
			log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
			log.SetOutput(os.Stderr)
		} else {
			log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
			log.SetOutput(io.MultiWriter(file, os.Stderr))
		}
		return file
	case "close":
		log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
		log.SetOutput(io.MultiWriter())
		return nil
	default:
		return nil
	}

}
func DetectionLogSize(file *os.File, size int64, path string, name string, writer string) {

	if file != nil {
		Ticker := time.NewTicker(time.Duration(1) * time.Second) //单位分钟		1
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

}
func DetectionLogModTime(file *os.File, path string, name string, day int64) {
	if file != nil {
		Ticker := time.NewTicker(1 * time.Second) //小时
		defer Ticker.Stop()
		clearLogForwardUnix := 60 * 60 * 24 * day
		for {
			select {
			case <-Ticker.C:
				Ticker.Reset(1 * time.Hour)
				fileInfoList, _ := ioutil.ReadDir(path)
				for i := range fileInfoList {
					// 判断文件是否空的
					if fileInfoList[i].Size() == 64 || (time.Now().Unix()-clearLogForwardUnix) > fileInfoList[i].ModTime().Unix() {
						if strings.Contains(fileInfoList[i].Name(), name) == true && strings.Contains(fileInfoList[i].Name(), ".log") == true && (path+"/"+fileInfoList[i].Name() != file.Name()) {
							delFile(path+"/"+fileInfoList[i].Name(), fileInfoList[i].IsDir())
						}
					}
					runtime.Gosched()
				}
			}
			runtime.Gosched()
		}
	}

}

// 删除文件
func delFile(str string, flag bool) {
	log.Println("delFile ClearingExpiredLogs ---- START", str)
	// 文件夹
	if flag {
		_ = os.RemoveAll(str)
	} else {
		_ = os.Remove(str)
	}
}
