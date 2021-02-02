// struct.go
package util_db

type StrokeAnalysisStruct struct {
	Code       []int     //代码
	Value      []float64 //变量
	Offset     []float64 //偏移量
	Directtion []int     //方向
	Abrupt     []int     //突变
	Gradual    []int     //渐变
	Continue   []bool    //延用
}
type SPMAnalysisStruct struct {
	Code       []int     //代码
	Value      []float64 //变量
	Offset     []float64 //偏移量
	Abrupt     []int     //突变
	Gradual    []int     //渐变
	Directtion []int     //方向
	Return     []int     //返回
}
type AreaAnalysisStruct struct {
	Code       []int     //代码
	Value      []float64 //变量
	Offset     []float64 //偏移量
	Abrupt     []int     //突变
	Gradual    []int     //渐变
	Directtion []int     //方向
	Return     []int     //返回
}
type FullnessCoefficientAnalysisStruct struct {
	Code       []int     //代码
	Value      []float64 //变量
	Offset     []float64 //偏移量
	Abrupt     []int     //突变
	Gradual    []int     //渐变
	Directtion []int     //方向
	Return     []int     //返回
}
type TopLeakCoefficientAnalysisStruct struct {
	Code       []int     //代码
	Value      []float64 //变量
	Offset     []float64 //偏移量
	Abrupt     []int     //突变
	Gradual    []int     //渐变
	Directtion []int     //方向
	Return     []int     //返回
}
type BottomLeakCoefficientAnalysisStruct struct {
	Code       []int     //代码
	Value      []float64 //变量
	Offset     []float64 //偏移量
	Abrupt     []int     //突变
	Gradual    []int     //渐变
	Directtion []int     //方向
	Return     []int     //返回
}
