package goreloaded

import (
	"strings"
	"unicode"
)

// func Quote(text string) []string {
//     statut := false
//     var result []rune

//     slice := []rune(text)
//     for i := 0; i < len(slice); i++ {
//         if slice[i] == '\'' && (i+1 < len(slice) && i-1 >= 0) && (unicode.IsLetter(slice[i+1]) && unicode.IsLetter(slice[i-1])) {
//             result = append(result, rune(slice[i]))
//             continue
//         }
//         if slice[i] == '\'' {
//             if !statut {
//                 statut = true
//                 if i != len(slice)-1{

//                     if i-1 >= 0 && slice[i-1] != ' ' && len(result) > 0 && result[len(result)-1] != ' ' {
//                         result = append(result, ' ')
//                     }
//                 }
//                 result = append(result, rune(slice[i]))
//                 if i+1 < len(slice) && slice[i+1] == ' ' {
//                     i++
//                 }
//             } else {
//                 statut = false
//                 if len(result) > 0 && result[len(result)-1] == ' ' {
//                     result = result[:len(result)-1]
//                 }
//                 result = append(result, rune(slice[i]))
//                 if (i+1 < len(slice)) && (slice[i+1] != ' ' && slice[i+1] != '\n' && !Only(string(slice[i+1]))) {
//                     result = append(result, ' ')
//                 }
//             }
//         } else {
//             result = append(result, rune(slice[i]))
//         }
//     }
//     sliceRes := strings.Split(string(result), " ")

//     return sliceRes
// }

func QuoteFmt(text string) string {
	res := ""
	temp := ""
	runes := []rune(text)
	clfound := false
	for i := 0; i < len(runes); i++ {
		if runes[i] == '\'' {
			if i > 0 && i < len(runes)-1 {
				if (unicode.IsLetter(runes[i+1]) || unicode.IsDigit(runes[i+1])) && (unicode.IsLetter(runes[i-1]) || unicode.IsDigit(runes[i-1])) {
					res += string(runes[i])
					continue
				}
			}
			if i < len(runes)-1 {
				clfound = false
				for j := i + 1; j < len(runes); j++ {
					if runes[j] == '\'' {
						if j > 0 && j < len(runes)-1 && (unicode.IsLetter(runes[j-1]) || unicode.IsDigit(runes[j-1])) && (unicode.IsLetter(runes[j+1]) || unicode.IsDigit(runes[j+1])) {
							temp += "'"
							continue
						}
						clfound = true
						i = j
						break
					}
					temp += string(runes[j])
				}
				if clfound {
					temp = strings.TrimSpace(temp)
					res += "'" + temp + "'"
					if i+1 < len(runes) && (unicode.IsLetter(runes[i+1]) || unicode.IsDigit(runes[i+1]) || runes[i+1] == '\'') {
						res += " "
					}
					temp = ""
					continue
				}
			}
		}
		res += string(runes[i])
	}
	return res
}
