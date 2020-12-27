// time.go
package util

import (
	"strings"
	"time"
)

func GetTimeStamp(strTime string) int64 {
	ts, err := time.ParseInLocation("2006-01-02 15:04:05", strTime, time.Local)
	if err == nil {
		return ts.Unix()
	} else {
		return 0
	}
}

func CalcTimeRange(testTime, startTime, endTime string, DelayTime int) int {
	if startTime == "00:00" && endTime == "24:00" {
		return 2 //2-代表全天生产，当前测试时间在开始结束时间内
	} else if startTime == "" && endTime == "" {
		return 1 //1-代表全天计划停井
	} else if startTime != "" && endTime != "" {
		if len(strings.Split(startTime, ":")) == 2 {
			startTime += ":00"
		}
		if len(strings.Split(endTime, ":")) == 2 {
			endTime += ":00"
		}
		startDate := strings.Split(testTime, " ")[0]
		endDate := strings.Split(testTime, " ")[0]

		testTS := GetTimeStamp(testTime)
		startTS := GetTimeStamp(startDate+" "+startTime) + int64(DelayTime)
		/*结束时间特殊处理*/
		if endTime == "24:00:00" {
			endTime = "00:00:00"
			endTimeFormat, _ := time.ParseInLocation("2006-01-02 15:04:05", endDate+" "+endTime, time.Local)
			endTimeFormat.AddDate(0, 0, 1)
			endDate = strings.Split(endTimeFormat.AddDate(0, 0, 1).Format("2006-01-02 15:04:05"), " ")[0]
		}
		endTS := GetTimeStamp(endDate+" "+endTime) - int64(DelayTime)
		// fmt.Printf("[calcTimeRange] testTime: %s, startTime: %s, endTime: %s, startdate: %s, endDate: %s,testTS: %d, startTS: %d, endTS: %d, DelayTime: %d\n", testTime, startTime, endTime, startDate, endDate, testTS, startTS, endTS, DelayTime)
		if testTS >= startTS && testTS <= endTS {
			return 3 //代表间抽生产，当前测试时间在开始结束时间内
		} else {
			return 4 //4-代表间抽生产，当前测试时间不在开始结束时间内
		}
	}
	return -1
}

func CalcTimeRange2(testTime, startTime, endTime string, DelayTime int) int {
	if startTime == "00:00" && endTime == "24:00" {
		return 2 //2-代表全天生产，当前测试时间在开始结束时间内
	} else if startTime == "" && endTime == "" {
		return 1 //1-代表全天计划停井
	} else if startTime != "" && endTime != "" {
		if len(strings.Split(startTime, ":")) == 2 {
			startTime += ":00"
		}
		if len(strings.Split(endTime, ":")) == 2 {

			endTime += ":00"
		}
		startDate := strings.Split(testTime, " ")[0]
		endDate := strings.Split(testTime, " ")[0]

		//		lastTime, _ := time.ParseInLocation("2006-01-02 15:04:05", temp.Last.TestTime, time.Local)
		testTS := GetTimeStamp(testTime)
		startTS := GetTimeStamp(startDate+" "+startTime) - int64(DelayTime)
		/*结束时间特殊处理*/
		if endTime == "24:00:00" {
			endTime = "00:00:00"
			endTimeFormat, _ := time.ParseInLocation("2006-01-02 15:04:05", endDate+" "+endTime, time.Local)
			endTimeFormat.AddDate(0, 0, 1)
			endDate = strings.Split(endTimeFormat.AddDate(0, 0, 1).Format("2006-01-02 15:04:05"), " ")[0]
		}

		endTS := GetTimeStamp(endDate+" "+endTime) + int64(DelayTime)
		// fmt.Println("endTime:", endTime)
		// fmt.Println("endTS:", endTS)
		// fmt.Printf("[calcTimeRange2] testTime: %s, startTime: %s, endTime: %s, startdate: %s, endDate: %s,testTS: %d, startTS: %d, endTS: %d, DelayTime: %d\n", testTime, startTime, endTime, startDate, endDate, testTS, startTS, endTS, DelayTime)
		if testTS >= startTS && testTS <= endTS {
			return 3 //代表间抽生产，当前测试时间在开始结束时间内
		} else {
			return 4 //4-代表间抽生产，当前测试时间不在开始结束时间内
		}
	}
	return -1
}
