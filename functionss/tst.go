package goreloaded

import "strings"

func Tst(s string) string {
	return strings.TrimPrefix(s,"'")
}