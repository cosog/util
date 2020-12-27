// log
package util

import (
	"fmt"
	"log"
	"os"
)

var File *os.File
var Logger *log.Logger
var Err error

func LogOpen() {
	File, Err = os.OpenFile("AgileCalculation.log", os.O_WRONLY|os.O_APPEND, 0)
	if Err != nil {
		fmt.Println("fail to open AgileCalculation.log file!")
		File, Err = os.Create("AgileCalculation.log")
		if Err != nil {
			log.Fatalln("fail to create AgileCalculation.log file!")
		} else {
			log.Println("success to create AgileCalculation.log file!")
		}
	} else {
		fmt.Println("success to opent AgileCalculation.log file")
	}

	Logger = log.New(File, "", log.LstdFlags|log.Llongfile)
	//	logger.SetFlags(log.LstdFlags)

}
