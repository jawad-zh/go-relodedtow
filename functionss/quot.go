package goreloaded

import (
	"strings"
	"unicode"
)

func Quotes(text string) string {
	Result := ""
	temp := ""
	runes := []rune(text)
	check := false
	for i := 0; i < len(runes); i++ {
		if runes[i] == '\'' {
			if i > 0 && i < len(runes)-1 {
				if (unicode.IsLetter(runes[i+1]) || unicode.IsDigit(runes[i+1])) && (unicode.IsLetter(runes[i-1]) || unicode.IsDigit(runes[i-1])) {
					Result += string(runes[i])
					continue
				}
			}
			if i < len(runes)-1 {
				check = false
				for j := i + 1; j < len(runes); j++ {
					if runes[j] == '\'' {
						if j > 0 && j < len(runes)-1 && (unicode.IsLetter(runes[j-1]) || unicode.IsDigit(runes[j-1])) && (unicode.IsLetter(runes[j+1]) || unicode.IsDigit(runes[j+1])) {
							temp += "'"
							continue
						}
						check = true
						i = j
						break
					}
					temp += string(runes[j])
				}
				if check {
					temp = strings.TrimSpace(temp)
					if len(Result) > 0 && unicode.IsLetter(rune(Result[len(Result)-1])) {
						Result += " "
					}
					Result += "'" + temp + "'"
					if i+1 < len(runes) && (unicode.IsLetter(runes[i+1]) || unicode.IsDigit(runes[i+1]) || runes[i+1] == '\'') {
						Result += " "
					}
					temp = ""
					continue
				}

			}
		}
		Result += string(runes[i])
	}
	return Result
}
