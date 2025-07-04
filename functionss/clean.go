package goreloaded

func Clean(s string)[]string{
	var word string 
	var words []string
	runes:= []rune(s)
	for i:=0 ; i < len(runes) ; i++{
		char := runes[i]
		if char != ' ' && char != '\n' && char !='\r' && char != '\u200b' {
			word += string(char)
		}else{
			if word != ""{
				words = append(words, word)
				word= ""
			}
		}
		if char == '\n' || char == '\r' || char == '\u200b' {
			words = append(words, string(char))
		}
	}
	if word != ""{
		words = append(words, word)
	}
	var Result []string 
	for j:=0 ; j < len(words) ; j++{
		if words[j] != ""{
			Result = append(Result, words[j])
		}
	}
	return Result
}