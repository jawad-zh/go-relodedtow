package goreloaded

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// )

// func Convert(s []string) []string {
// 	for i := 0; i < len(s); i++ {
// 		// word := strings.ToLower(s[i])

// 		if i > 0 && (s[i] == "(up)") {
// 			for j := i - 1; j >= 0; j-- {
// 				if !IsControlStr(s[j]) {
// 					s[j] = strings.ToUpper(s[j])
// 					break
// 				}
// 			}
// 			s[i] = ""
// 			s = Rmov(s)
// 			i--
// 		} else if i == 0 && (s[i] == "(up)") {
// 			s[i] = ""
// 			s = Rmov(s)
// 			i--
// 		} else if i > 0 && (s[i] == "(low)") {
// 			for j := i - 1; j >= 0; j-- {
// 				if !IsControlStr(s[j]) {
// 					s[j] = strings.ToLower(s[j])
// 					break
// 				}
// 			}
// 			s[i] = ""
// 			s = Rmov(s)
// 			i--
// 		} else if i == 0 && (s[i] == "(low)") {
// 			s[i] = ""
// 			s = Rmov(s)
// 			i--
// 		} else if i > 0 && (s[i] == "(cap)") {
// 			for j := i - 1; j >= 0; j-- {
// 				if !IsControlStr(s[j]) {
// 					s[j] = Capitalize(s[j])
// 					break
// 				}
// 			}
// 			s[i] = ""
// 			s = Rmov(s)
// 			i--
// 		} else if i == 0 && (s[i] == "(cap)") {
// 			s[i] = ""
// 			s = Rmov(s)
// 			i--
// 		} else if i == 0 && (s[i] == "(hex)") {
// 			s[i] = ""
// 			s = Rmov(s)
// 			i--
// 		} else if i > 0 && (s[i] == "(hex)") {
// 			for j := i - 1; j >= 0; j-- {
// 				if !IsControlStr(s[j]) {
// 					decimal, err := strconv.ParseInt(s[j], 16, 64)
// 					if err != nil {
// 						fmt.Println("Error: invalid hex:", s[j])
// 					} else {
// 						s[j] = strconv.Itoa(int(decimal))
// 					}
// 					break
// 				}
// 			}
// 			s[i] = ""
// 			s = Rmov(s)
// 			i--
// 		} else if i > 0 && (s[i] == "(bin)") {
// 			for j := i - 1; j >= 0; j-- {
// 				if !IsControlStr(s[j]) {
// 					decimal, err := strconv.ParseInt(s[j], 16, 64)
// 					if err != nil {
// 						fmt.Println("Error: invalid hex:", s[j])
// 					} else {
// 						s[j] = strconv.Itoa(int(decimal))
// 					}
// 					break
// 				}
// 			}
// 			s[i] = ""
// 			s = Rmov(s)
// 			i--
// 		} else if i == 0 && (s[i] == "(bin)") {
// 			s[i] = ""
// 			s = Rmov(s)
// 			i--
// 		}
// 		if IsFlag(s) {
// 			if i != len(s)-1 && CorrectEnd(s[i+1]) {
// 				if s[i] == "(up," && i+1 < len(s) {
// 					nStr := strings.TrimSuffix(s[i+1], ")")
// 					n, err := strconv.Atoi(nStr)
// 					if err != nil {
// 						fmt.Println("Error Case:", err)
// 						continue
// 					}
// 					count := 0
// 					for j := i - 1; j >= 0 && count < n; j-- {
// 						if !IsControlStr(s[j]) {
// 							s[j] = strings.ToUpper(s[j])
// 							count++
// 						}
// 					}
// 				}
// 				s[i] = ""
// 				s[i+1] = ""
// 				s = Rmov(s)
// 				i--
// 			} else if s[i] == "(low," && i+1 < len(s) {
// 				nStr := strings.TrimSuffix(s[i+1], ")")
// 				n, err := strconv.Atoi(nStr)
// 				if err != nil {
// 					fmt.Println("Error Case:", err)
// 					continue
// 				}
// 				count := 0
// 				for j := i - 1; j >= 0 && count < n; j-- {
// 					if !IsControlStr(s[j]) {
// 						s[j] = strings.ToLower(s[j])
// 						count++
// 					}
// 				}
// 				s[i] = ""
// 				s[i+1] = ""
// 				s = Rmov(s)
// 				i--
// 			} else if s[i] == "(cap," && i+1 < len(s) {
// 				nStr := strings.TrimSuffix(s[i+1], ")")
// 				n, err := strconv.Atoi(nStr)
// 				if err != nil {
// 					fmt.Println("Error Case:", err)
// 					continue
// 				}
// 				count := 0
// 				for j := i - 1; j >= 0 && count < n; j-- {
// 					if !IsControlStr(s[j]) {
// 						s[j] = Capitalize(s[j])
// 						count++
// 					}
// 				}
// 				s[i] = ""
// 				s[i+1] = ""
// 				s = Rmov(s)
// 				i--

// 			}
// 		}

// 	}
// 	return s
// }

func Rmov(s []string) []string {

	var result []string
	for _, word := range s {
		if word != "" {
			result = append(result, word)
		}
	}
	return result
}
