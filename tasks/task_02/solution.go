package main

import (
	"unicode/utf8"
)

func rotateRunes(s string, shift int) string {
	if s == "" {
		return ""
	}

	r := []rune(s)

	//если сдвига нет, битые руны всё равно нужно заменить.
	if shift == 0 {
		return string(r)
	}

	//проверяем наше число сдвига, чтобы не выходило за пределы длины строки
	check_shift := shift % utf8.RuneCountInString(s)
	if shift < 0 {
		check_shift += utf8.RuneCountInString(s)
	}

	result := append(r[check_shift:], r[:check_shift]...)
	return string(result)

}
