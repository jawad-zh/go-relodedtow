package goreloaded

import (
    "unicode"
)

func Capitalize(s string) string {
    slice := []rune(s)
    flag := false
    
    for i := 0; i < len(slice); i++ {
        if   !flag {
            slice[i] = unicode.ToUpper(slice[i])
            flag = true
        } else if flag {
            slice[i] = unicode.ToLower(slice[i])
        }
    }
    return string(slice)
}