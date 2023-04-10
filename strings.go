package utils

import (
	"strconv"
	"strings"
)

func HasDigit(s string) bool {
	b := strings.IndexFunc(s, func(c rune) bool { return c >= '0' && c <= '9' }) >= 0
	return b
}

func HasNotDigit(s string) bool {
	_, err := strconv.Atoi(s)
	return err != nil
}

func OnlyDigits(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func SameStringSlice(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	diff := make(map[string]int, len(x))
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
