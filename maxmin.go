// maxmin.go
package util

import (
	"math"
	"runtime"
)

func MaxValue(vals ...float64) float64 {
	max := -math.MaxFloat64
	for _, v := range vals {
		if v > max {
			max = v
		}
		runtime.Gosched()
	}
	return max
}

func MinValue(vals ...float64) float64 {
	min := math.MaxFloat64
	for _, v := range vals {
		if v < min {
			min = v
		}
		runtime.Gosched()
	}
	return min
}
