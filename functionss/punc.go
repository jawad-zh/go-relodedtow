package goreloaded

func Punc(s []string) []string {
	var result []string
	i := 0
	for i < len(s) {
		if Only(s[i]) {
			group := s[i]
			j := i + 1
			for j < len(s) && Only(s[j]) {
				group += s[j]
				j++
			}
			result = append(result, group)
			i = j
		} else {
			result = append(result, s[i])
			i++
		}
	}
	for j := 0; j < len(result); j++ {
		if Only(result[j]) {
			if j != 0 {
				result[j-1] = result[j-1] + result[j]
				result[j] = ""
			}
		}
	}
	var Result []string
	for _, r := range result {
		if r != "" {
			Result = append(Result, string(r))
		}
	}
	// TokenizeSmarte(Result)
	return Result
}

// func isPunc(r rune) bool {
// 	return strings.ContainsRune(".,!?;:", r)
// }

// func TokenizeSmarte(words []string) []string {
// 	var result []string

// 	for _, word := range words {
// 		runes := []rune(word)
// 		var token []rune

// 		i := 0
// 		for i < len(runes) {
// 			r := runes[i]

// 			if isPunc(r) {
// 				if i > 0 && i < len(runes)-1 &&
// 					unicode.IsLetter(runes[i-1]) && unicode.IsLetter(runes[i+1]) {
// 					token = append(token, r)
// 					i++
// 					continue
// 				}
// 				if len(token) > 0 {
// 					result = append(result, string(token))
// 					token = []rune{}
// 				}
// 				puncGroup := []rune{r}
// 				j := i + 1
// 				for j < len(runes) && isPunc(runes[j]) {
// 					puncGroup = append(puncGroup, runes[j])
// 					j++
// 				}
// 				result = append(result, string(puncGroup))
// 				i = j
// 			} else {
// 				token = append(token, r)
// 				i++
// 			}
// 		}

// 		if len(token) > 0 {
// 			result = append(result, string(token))
// 		}
// 	}

// 	return result
// }
