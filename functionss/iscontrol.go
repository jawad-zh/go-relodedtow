package goreloaded

import (
	"unicode"
	"unicode/utf8"
)

func IsControlStr(s string) bool {
	r, _ := utf8.DecodeRuneInString(s)
	return unicode.IsControl(r)
}





