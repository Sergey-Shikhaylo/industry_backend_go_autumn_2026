package main

import (
	"strings"
	"unicode/utf8"
)

func runeIdx(s string, number int) int {
	count := 1
	for i := range s {
		if count == number {
			return i
		}
		count++
	}
	return 0
}

func rotateRunes(s string, shift int) string {
	if s == "" {
		return ""
	}

	if shift == 0 {
		return s
	}

	s = strings.ToValidUTF8(s, "")
	check_shift := shift
	if shift > utf8.RuneCountInString(s) {
		check_shift = shift % utf8.RuneCountInString(s)
	}
	shiftIdx := runeIdx(s, check_shift)

	return s[shiftIdx:] + s[:shiftIdx]

}
