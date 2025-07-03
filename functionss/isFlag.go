package goreloaded

import "strings"

func CorrectEnd(s string) bool {
	for i := 0; i < len(s); i++ {

		if i+1 < len(s) && strings.HasSuffix(string(s[i+1]), ")") {
			return true
		}
	}
	return false
}
func IsFlag(s []string) bool {
	for i := 0; i < len(s); i++ {
		if (s[i] == "(up," && i+1 < len(s)) && CorrectEnd(s[i+1]) {
			return true
		}else if (s[i] == "(low," && i+1 < len(s)) && CorrectEnd(s[i+1]) {
			return true
		}else if (s[i] == "(cap," && i+1 < len(s)) && CorrectEnd(s[i+1]) {
			return true
		}else if (s[i] == "(hex," && i+1 < len(s)) && CorrectEnd(s[i+1]){
			return true
		}else if (s[i] == "(bin," && i+1 < len(s)) && CorrectEnd(s[i+1]){
			return true
		}
	}
	return false
}
