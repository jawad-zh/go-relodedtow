package goreloaded

import (
	"strings"
	"unicode"
)

func Quote(text string) []string {
    statut := false
    var result []rune
    
    slice := []rune(text)
    for i := 0; i < len(slice); i++ {
        if slice[i] == '\'' && (i+1 < len(slice) && i-1 >= 0) && (unicode.IsLetter(slice[i+1]) && unicode.IsLetter(slice[i-1])) {
            result = append(result, rune(slice[i]))
            // i++
            continue
        }
        if slice[i] == '\'' {
            if !statut {
                statut = true
                if i-1 >= 0 && slice[i-1] != ' ' && len(result) > 0 && result[len(result)-1] != ' ' {
                    result = append(result, ' ')
                }
                result = append(result, rune(slice[i]))
                if i+1 < len(slice) && slice[i+1] == ' ' {
                    i++
                }
            } else {
                statut = false
                if len(result) > 0 && result[len(result)-1] == ' ' {
                    result = result[:len(result)-1]
                }
                result = append(result, rune(slice[i]))
                if (i+1 < len(slice)) && (slice[i+1] != ' ' && slice[i+1] != '\n' && !Only(string(slice[i+1]))) {
                    result = append(result, ' ')
                }
            }
        } else {
            result = append(result, rune(slice[i]))
        }
    }
    sliceRes := strings.Split(string(result), " ")

    return sliceRes
}