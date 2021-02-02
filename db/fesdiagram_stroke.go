// fesdiagram_stroke.go
package util_db

import (
	"sync"

	"github.com/cosog/util"
	"github.com/cosog/util/slice"
)

func StrokeHistoryAnalysis(m map[string]StrokeAnalysisStruct, wellName string, code int, stroke float64, length int) int {

	var rwMutex sync.RWMutex

	rwMutex.RLock() //读锁定

	v, _ := m[wellName]

	rwMutex.RUnlock() //读解锁

	//1-追加、截取
	v.Value = append(v.Value, stroke)
	util_slice.Float64LengthCut(&v.Value, length)
	cnt := len(v.Value)

	//2-填充偏移量/计算面积突变、渐变与不变
	if cnt > 1 {

		v.Offset = append(v.Offset, v.Value[cnt-1]-v.Value[cnt-2])

		if (v.Value[cnt-1] - v.Value[cnt-2]) > util.CalculationError {

			v.Directtion = append(v.Directtion, 1)

			if v.Value[cnt-1] > 2*v.Value[cnt-2] {

				//增大突变
				v.Abrupt = append(v.Abrupt, 1)
				v.Gradual = append(v.Gradual, 0)

			} else {

				//增大渐变
				v.Abrupt = append(v.Abrupt, 0)
				v.Gradual = append(v.Gradual, 1)

			}

		} else if (v.Value[cnt-2] - v.Value[cnt-1]) > util.CalculationError {

			v.Directtion = append(v.Directtion, -1)

			if 2*v.Value[cnt-1] < v.Value[cnt-2] {
				//减小突变
				v.Abrupt = append(v.Abrupt, -1)

				//加强采集异常判断
				// if MIValueSlice.

				v.Gradual = append(v.Gradual, 0)
			} else {
				//减小渐变
				v.Abrupt = append(v.Abrupt, 0)
				v.Gradual = append(v.Gradual, -1)
			}

		} else { //初始数据对齐，填写默认值

			v.Directtion = append(v.Directtion, 0)
			v.Abrupt = append(v.Abrupt, 0)
			v.Gradual = append(v.Gradual, 0)
		}
	} else {
		v.Directtion = append(v.Directtion, 0)
		v.Abrupt = append(v.Abrupt, 0)
		v.Gradual = append(v.Gradual, 0)
	}

	util_slice.Float64LengthCut(&v.Offset, length)
	util_slice.IntLengthCut(&v.Directtion, length)
	util_slice.IntLengthCut(&v.Abrupt, length)
	util_slice.IntLengthCut(&v.Gradual, length)

	rwMutex.Lock() //写锁定
	m[wellName] = v
	rwMutex.Unlock() //写解锁

	switch code { //如果是条带状
	case 1202:
		fallthrough //游动凡尔失灵/油管漏
	case 1203:
		fallthrough //固定凡尔失灵
	case 1204: //双凡尔失灵
		fallthrough
	case 1205:
		fallthrough //杆断脱
	case 1206: //柱塞未下入工作筒
		fallthrough
	case 1207: //抽喷

		if cnt > 1 {
			if v.Abrupt[cnt-1] == 1 || v.Abrupt[cnt-1] == -1 || v.Continue[cnt-1] == true {
				v.Continue = append(v.Continue, true)
				util_slice.BoolLengthCut(&v.Continue, length)
				return 1232
			} else {
				v.Continue = append(v.Continue, false)
				util_slice.BoolLengthCut(&v.Continue, length)
				return code
			}

		} else {
			v.Continue = append(v.Continue, false)
			util_slice.BoolLengthCut(&v.Continue, length)
			return code

		}

	default:
		v.Continue = append(v.Continue, false)
		util_slice.BoolLengthCut(&v.Continue, length)
		return code
	}
}
