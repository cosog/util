// fesdiagram_area.go
package util_db

// import (
// 	"github.com/cosog/util"
// 	"github.com/cosog/util/slice"
// )

func AreaHistoryAnalysis(m map[string]AreaAnalysisStruct, wellName string, code int, area float64, length int) {

	// for {
	// 	select {
	// 	case <-FESDiagramHistoryAreaChan.Start:
	// 		cnt := len(s.Area)
	// 		//1-追加、截取
	// 		s.Area = append(s.Area, h.Area)
	// 		util_slice.Float64LengthCut(&s.Area, length)

	// 		//2-填充偏移量/计算面积突变、渐变与不变
	// 		if cnt > 1 {

	// 			g.Area.Offset = append(g.Area.Offset, s.Area[cnt-1]-s.Area[cnt-2])

	// 			if (g.Area.Offset[cnt-1] - g.Area.Offset[cnt-2]) > util.CalculationError {

	// 				g.Area.Directtion = append(g.Area.Directtion, 1)

	// 				if g.Area.Offset[cnt-1] > 2*g.Area.Offset[cnt-2] {

	// 					//增大突变
	// 					g.Area.Abrupt = append(g.Area.Abrupt, 1)
	// 					g.Area.Gradual = append(g.Area.Gradual, 0)

	// 				} else {

	// 					//增大渐变
	// 					g.Area.Abrupt = append(g.Area.Abrupt, 0)

	// 					switch h.ResultCode { //如果是条带状
	// 					case 1202:
	// 						fallthrough //游动凡尔失灵/油管漏
	// 					case 1203:
	// 						fallthrough //固定凡尔失灵
	// 					case 1204: //双凡尔失灵
	// 						fallthrough
	// 					case 1205:
	// 						fallthrough //杆断脱
	// 					case 1206: //柱塞未下入工作筒
	// 						fallthrough
	// 					case 1207: //抽喷

	// 					}
	// 					g.Area.Gradual = append(g.Area.Gradual, 1)

	// 				}

	// 			} else if (g.Area.Offset[cnt-2] - g.Area.Offset[cnt-1]) > util.CalculationError {

	// 				g.Area.Directtion = append(g.Area.Directtion, -1)

	// 				if 2*g.Area.Offset[cnt-1] < g.Area.Offset[cnt-2] {
	// 					//减小突变
	// 					g.Area.Abrupt = append(g.Area.Abrupt, -1)

	// 					//加强采集异常判断
	// 					// if MIValueSlice.

	// 					g.Area.Gradual = append(g.Area.Gradual, 0)
	// 				} else {
	// 					//减小渐变
	// 					g.Area.Abrupt = append(g.Area.Abrupt, 0)
	// 					g.Area.Gradual = append(g.Area.Gradual, -1)
	// 				}

	// 			} else {
	// 				//不变
	// 				g.Area.Directtion = append(g.Area.Directtion, 0)
	// 				g.Area.Abrupt = append(g.Area.Abrupt, 0)
	// 				g.Area.Gradual = append(g.Area.Gradual, 0)
	// 			}
	// 		}
	// 		util_slice.IntLengthCut(&g.Area.Directtion, length)
	// 		util_slice.IntLengthCut(&g.Area.Abrupt, length)
	// 		util_slice.IntLengthCut(&g.Area.Gradual, length)
	// 		//6-

	// 		if g.Area.Directtion[cnt-1] != 0 {
	// 			//
	// 		}
	// 		goto Loop
	// 	}
	// 	runtime.Gosched()
	// }
	// Loop:
	// FESDiagramHistoryAreaChan.End <- true
}
