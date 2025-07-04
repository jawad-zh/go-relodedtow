package goreloaded

func SplitPunc(s string) []string {
	var word string
	var words []string
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' && !Only(string(s[i])) {
			word += string(s[i])
		} else {
			if word != "" {
				words = append(words, word)
				word = ""
			}
			if Only(string(s[i])) {
				words = append(words, string(s[i]))
			}
		}
	}
	if word != ""{
		words = append(words,word)
	}
	return words
}
