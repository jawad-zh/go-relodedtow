package goreloaded


func JoinCleaned(words []string) string {
	var result string
	for i, w := range words {
		if w == "\n" || w == "\r" {
			result += w
		} else {
			if i > 0 && words[i-1] != "\n" && words[i-1] != "\r" {
				result += " "
			}
			result += w
		}
	}
	return result
}
