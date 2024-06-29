// fesinterpolation.go
package util_fes

import (
	"math"

	"github.com/cosog/util"
)

func FESInterpolationPoint(fes *FESStruct, targetCNT int) FESStruct {

	var tempFES = new(FESStruct)
	var resFES = new(FESStruct)
	var x float64
	var start, end int
	var sMin float64 = math.MaxFloat64
	var sMax float64 = -math.MaxFloat64

	if targetCNT != 0 {
		x = float64(fes.CNT) / float64(targetCNT)

		for i := 0; i < targetCNT; i++ {
			start = int(float64(i) * x)
			// if start == targetCNT-1 {
			if start == fes.CNT-1 {
				end = 0
			} else {
				end = start + 1
			}
			if len(fes.S) > 0 {
				tempFES.S = append(tempFES.S, fes.S[start]+(float64(i)*x-float64(start))*(fes.S[end]-fes.S[start]))
				if len(fes.F) != 0 {
					tempFES.F = append(tempFES.F, fes.F[start]+(float64(i)*x-float64(start))*(fes.F[end]-fes.F[start]))
				}
				if len(fes.Watt) != 0 {
					tempFES.Watt = append(tempFES.Watt, fes.Watt[start]+(float64(i)*x-float64(start))*(fes.Watt[end]-fes.Watt[start]))
				}
				if len(fes.I) != 0 {
					tempFES.I = append(tempFES.I, fes.I[start]+(float64(i)*x-float64(start))*(fes.I[end]-fes.I[start]))
				}
			}
		}
		//拉伸
		for _, v := range tempFES.S {
			if sMax < v {
				sMax = v
			}
			if sMin > v {
				sMin = v
			}
		}
		if math.Abs(sMax-sMin) < util.CalculationError {
			resFES.S = append(resFES.S, tempFES.S...)

		} else {
			for _, v := range tempFES.S {
				resFES.S = append(resFES.S, sMin+(v-sMin)/(sMax-sMin)*fes.Stroke)
			}
		}

		resFES.F = append(resFES.F, tempFES.F...)
		resFES.Watt = append(resFES.Watt, tempFES.Watt...)
		resFES.I = append(resFES.I, tempFES.I...)

		resFES.CNT = targetCNT
		resFES.AcqTime = fes.AcqTime
		resFES.SPM = fes.SPM
		resFES.Stroke = fes.Stroke

	}

	return *resFES
}
