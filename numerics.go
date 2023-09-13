package utils

import "math"

func SameIntSlice(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	diff := make(map[int]int, len(x))
	for _, _x := range x {
		diff[_x]++
	}
	for _, _y := range y {
		if _, ok := diff[_y]; !ok {
			return false
		}
		diff[_y] -= 1
		if diff[_y] == 0 {
			delete(diff, _y)
		}
	}
	return len(diff) == 0
}

func RoundF64(n float64) float64 {
	return math.Round(n*100) / 100
}
