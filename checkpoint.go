// util
package util

func CheckPoint(index int, CNT int) int {
	if index < 0 {
		index = CNT - 1
	}
	if index >= CNT {
		index = 0
	}

	return index
}
