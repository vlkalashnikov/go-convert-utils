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
