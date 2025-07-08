package goreloaded

import (
	"fmt"
	"strconv"
	"strings"
)

func Convert(s []string) []string {
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case "(up)", "(low)", "(cap)":
			if i > 0 {
				s[i-1] = applySingleFlag(s[i], s[i-1])
			}
			s[i] = ""
			s = Rmov(s)
			i=-1
		case "(hex)":
			if i > 0 {
				val, err := strconv.ParseInt(s[i-1], 16, 64)
				if err != nil {
					fmt.Println("Error: invalid hex:", s[i-1])
				} else {
					s[i-1] = strconv.Itoa(int(val))
				}
			}
			s[i] = ""
			s = Rmov(s)
			i=-1
		case "(bin)":
			if i > 0 {
				val, err := strconv.ParseInt(s[i-1], 2, 64)
				if err != nil {
					fmt.Println("Error: invalid binary:", s[i-1])
				} else {
					s[i-1] = strconv.Itoa(int(val))
				}
			}
			s[i] = ""
			s = Rmov(s)
			i=-1
		}
		// Handle (up,n), (low,n), (cap,n)
		if i < len(s)-1 && IsFlag(s) {
			nStr := strings.TrimSuffix(s[i+1], ")")
			n, err := strconv.Atoi(nStr)
			if err != nil {
				continue
			}
			switch s[i] {
			case "(up,":
				ApplyNFlag(s, i, n, strings.ToUpper)
			case "(low,":
				ApplyNFlag(s, i, n, strings.ToLower)
			case "(cap,":
				ApplyNFlag(s, i, n, Capitalize)
			}
			s[i] = ""
			s[i+1] = ""
			s = Rmov(s)
			i = -1

		}
	}
	return s
}
// apply the valid flag
func applySingleFlag(flag, word string) string {
	switch flag {
	case "(up)":
		return strings.ToUpper(word)
	case "(low)":
		return strings.ToLower(word)
	case "(cap)":
		return Capitalize(word)
	default:
		return word
	}
}
// apply the valid flag with n
func ApplyNFlag(s []string, index int, n int, transform func(string) string) {
	count := 0
	for j := index - 1; j >= 0 && count < n; j-- {
		s[j] = transform(s[j])
		count++
	}
}
// remove the empty string
func Rmov(s []string) []string {
	var result []string
	for _, word := range s {
		if word != "" {
			result = append(result, word)
		}
	}
	return result
}
