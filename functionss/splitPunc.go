package goreloaded

// append each punc in slice
func SplitPunc(s string) []string {
	var word string
	var words []string
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		char := runes[i]
		if char != ' ' && !Only(string(char))  {
			word += string(char)
		} else {
			if word != "" {
				words = append(words, word)
				word = ""
			}
			if Only(string(char)) {
				words = append(words, string(char))
			}
		}
	}
	if word != "" {
		words = append(words, word)
	}
	return words
}
