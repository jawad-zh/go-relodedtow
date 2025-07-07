package goreloaded

import (
	"strings"
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


func Quote(reloaded []string) []string {
    s := strings.Join(reloaded, " ")
    var open bool
    index := 0
    var str string
    for i, c := range s {
        if i != len(s)-1 && i != 0 {
            if !open {
                if c == '\'' && (s[i+1] == ' ' || s[i-1] == ' ' || s[i+1] == '\'' || s[i-1] == '\'' || Only(string(s[i+1])) || Only(string(s[i-1]))) {
                    str += s[index:i]
                    open = true
                    index = i
                }
            } else if open && c == '\'' && (s[i+1] == ' ' || s[i-1] == ' ' || s[i+1] == '\'' || s[i-1] == '\'' || Only(string(s[i+1])) || Only(string(s[i-1]))){
                if Only(string(s[i+1])) || Only(string(s[i-1])) {
                    str += " " + "'" + strings.TrimSpace(s[index+1:i]) + "'"
                    open = false
                    index = i + 1
                    continue
                }
                str += " " + "'" + strings.TrimSpace(s[index+1:i]) + "'" + " "
                open = false
                index = i + 1
                continue
            }
        } else if i == 0 {
            if c == '\'' {
                str += s[index:i]
                open = true
                index = i
            }
        }
        if i == len(s)-1 {
            if open && c == '\'' {
                str += " " + "'" + strings.TrimSpace(s[index+1:len(s)-1]) + "'"
            } else {
                str += s[index:]
            }
        }
    }
    return strings.Fields(str)
}