// struct.go
package util_fes

type FESStruct struct {
	AcqTime string    //采集时间
	Stroke  float64   //功图冲程			m
	SPM     float64   //功图冲次			1/min
	CNT     int       //点数
	F       []float64 //
	Watt    []float64 //
	I       []float64 //
	S       []float64 //
}
