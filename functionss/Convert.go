package goreloaded

import (
	"fmt"
	"strconv"
	"strings"
)

func Convert(s []string) []string {
	for i := 0; i < len(s); i++ {
		// word := strings.ToLower(s[i])
        
			if i > 0 && (s[i] == "(up)") {
			s[i-1] = strings.ToUpper(s[i-1])
			s[i] = ""
			s = Rmov(s)
			i--
		} else if i == 0 && (s[i] == "(up)") {
			s[i] = ""
			s = Rmov(s)
			i--
		} else if i > 0 && (s[i] == "(low)") {
			s[i-1] = strings.ToLower(s[i-1])
			s[i] = ""
			s = Rmov(s)
			i--
		} else if i == 0 && (s[i] == "(low)") {
			s[i] = ""
			s = Rmov(s)
			i--
		} else if i > 0 && (s[i] == "(cap)") {
			s[i-1] = Capitalize(s[i-1])
			s[i] = ""
			s = Rmov(s)
			i--
		} else if i == 0 && (s[i] == "(cap)") {
			s[i] = ""
			s = Rmov(s)
			i--
		} else if i == 0 && (s[i] == "(hex)") {
			s[i] = ""
			s = Rmov(s)
			i--
		} else if i > 0 && (s[i] == "(hex)") {
			decimal, err := strconv.ParseInt(s[i-1], 16, 64)
			if err != nil {
				fmt.Println("Error Case:", err)
			} else {
				s[i-1] = strconv.Itoa(int(decimal))
			}
			s[i] = ""
			s = Rmov(s)
			i--
		} else if i > 0 && (s[i] == "(bin)") {
			decimal, err := strconv.ParseInt(s[i-1], 2, 64)
			if err != nil {
				fmt.Println("Error Case:", err)
			} else {
				s[i-1] = strconv.Itoa(int(decimal))
			}
			s[i] = ""
			s = Rmov(s)
			i--
		} else if i == 0 && (s[i] == "(bin)") {
			s[i] = ""
			s = Rmov(s)
			i--
		}
		if IsFlag(s) {
			if i != len(s)-1 && CorrectEnd(s[i+1]) {
				if s[i] == "(up," && i+1 < len(s) {
					nStr := strings.TrimSuffix(s[i+1], ")")
					n, err := strconv.Atoi(nStr)
					if err != nil {
						fmt.Println("Error Case:", err)
						continue
					}
					for j := 0; j <= n; j++ {
						if i-j >= 0 {
							s[i-j] = strings.ToUpper(s[i-j])
						}
					}
					s[i] = ""
					s[i+1] = ""
					s = Rmov(s)
					i--
				} else if s[i] == "(low," && i+1 < len(s) {
					nStr := strings.TrimSuffix(s[i+1], ")")
					n, err := strconv.Atoi(nStr)
					if err != nil {
						fmt.Println("Error Case:", err)
						continue
					}
					for j := 0; j <= n; j++ {
						if i-j >= 0 {
							s[i-j] = strings.ToLower(s[i-j])
						}
					}
					s[i] = ""
					s[i+1] = ""
					s = Rmov(s)
					i--
				} else if s[i] == "(cap," && i+1 < len(s) {
					nStr := strings.TrimSuffix(s[i+1], ")")
					n, err := strconv.Atoi(nStr)
					if err != nil {
						fmt.Println("Error Case:", err)
						continue
					}
					for j := 0; j <= n; j++ {
						if i-j >= 0 {
							s[i-j] = Capitalize(s[i-j])
						}
					}
					s[i] = ""
					s[i+1] = ""
					s = Rmov(s)
					i--
				} else if s[i] == "(hex," && i+1 < len(s) {
					nStr := strings.TrimSuffix(s[i+1], ")")
					n, err := strconv.Atoi(nStr)
					if err != nil {
						fmt.Println("Error Case:", err)
						continue
					}
					for j := 0; j <= n; j++ {
						if i-j >= 0 {
							n, err := strconv.ParseInt(string(s[i-j]), 16, 64)
							if err != nil {
								fmt.Println("Error Case:", err)
							} else {
								s[i-j] = strconv.Itoa(int(n))
							}
						}
					}
					s[i] = ""
					s[i+1] = ""
					s = Rmov(s)
					i--
				} else if s[i] == "(bin," && i+1 < len(s) {
					nStr := strings.TrimSuffix(s[i+1], ")")
					n, err := strconv.Atoi(nStr)
					if err != nil {
						fmt.Println("Error Case:", err)
						continue
					}
					for j := 0; j <= n; j++ {
						if i-j >= 0 {
							n, err := strconv.ParseInt(string(s[i-j]), 2, 64)
							if err != nil {
								fmt.Println("Error Case:", err)
							} else {
								s[i-j] = strconv.Itoa(int(n))
							}
						}
					}
					s[i] = ""
					s[i+1] = ""
					s = Rmov(s)
					i--
				}
			}

		}
	}
		
	return s
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
