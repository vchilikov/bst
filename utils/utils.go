package utils

import "strings"

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Pad(s string, length int) string {
	if len(s) >= length {
		return s
	}
	pLen := length - len(s)
	return strings.Repeat(" ", (pLen+1)/2) + s + strings.Repeat(" ", pLen-(pLen+1)/2)
}
