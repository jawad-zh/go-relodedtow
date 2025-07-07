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
				for j := i - 1; j >= 0; j-- {
					if s[j] == "\n" {
						break
					}
					if !IsControlStr(s[j]) {
						s[j] = applySingleFlag(s[i], s[j])
						break
					}
				}
			}
			s[i] = ""
			s = Rmov(s)
			i--

		case "(hex)":
			if i > 0 {
				for j := i - 1; j >= 0; j-- {
					if s[j] == "\n" {
						break
					}
					if !IsControlStr(s[j]) {
						val, err := strconv.ParseInt(s[j], 16, 64)
						if err != nil {
							fmt.Println("Error: invalid hex:", s[j])
						} else {
							s[j] = strconv.Itoa(int(val))
						}
						break
					}
				}
			}
			s[i] = ""
			s = Rmov(s)
			i--

		case "(bin)":
			if i > 0 {
				for j := i - 1; j >= 0; j-- {
					if s[j] == "\n" {
						break
					}
					if !IsControlStr(s[j]) {
						val, err := strconv.ParseInt(s[j], 2, 64)
						if err != nil {
							fmt.Println("Error: invalid binary:", s[j])
						} else {
							s[j] = strconv.Itoa(int(val))
						}
						break
					}
				}
			}
			s[i] = ""
			s = Rmov(s)
			i--
		}

		// Handle (up,n), (low,n), (cap,n)
		if i < len(s)-1 && IsFlag(s) && CorrectEnd(s[i+1]) {
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
			i=-1

		}
	}
	return s
}

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

func ApplyNFlag(s []string, index int, n int, transform func(string) string) {
	count := 0
	for j := index - 1; j >= 0 && count < n; j-- {
		if s[j] == "\n" {
			break
		}
		if !IsControlStr(s[j]) {
			s[j] = transform(s[j])
			count++
		}
	}
}

func Rmov(s []string) []string {

	var result []string
	for _, word := range s {
		if word != "" {
			result = append(result, word)
		}
	}
	return result
}
