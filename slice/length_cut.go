// length_cut.go
package util_slice

import (
	"time"
)

func Float64LengthCut(v *[]float64, length int) {

	if len(*v) > length {
		*v = append((*v)[:0], (*v)[len(*v)-length:]...)
	}
}
func TimeLengthCut(v *[]time.Time, length int) {

	if len(*v) > length {
		*v = append((*v)[:0], (*v)[len(*v)-length:]...)
	}
}
func BoolLengthCut(v *[]bool, length int) {

	if len(*v) > length {
		*v = append((*v)[:0], (*v)[len(*v)-length:]...)
	}
}
func StringLengthCut(v *[]string, length int) {

	if len(*v) > length {
		*v = append((*v)[:0], (*v)[len(*v)-length:]...)
	}
}
func IntLengthCut(v *[]int, length int) {

	if len(*v) > length {
		*v = append((*v)[:0], (*v)[len(*v)-length:]...)
	}
}
