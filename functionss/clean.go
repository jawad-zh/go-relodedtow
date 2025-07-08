package goreloaded

func Clean(s string)[]string{
	var word string 
	var words []string
	runes:= []rune(s)
	for i:=0 ; i < len(runes) ; i++{
		char := runes[i]
		if char != ' ' {
			word += string(char)
		}else{
			if word != ""{
				words = append(words, word)
				word= ""
			}
		}
	}
	// last word
	if word != ""{
		words = append(words, word)
	}
	// clean ""
	var Result []string 
	for j:=0 ; j < len(words) ; j++{
		if words[j] != ""{
			Result = append(Result, words[j])
		}
	}
	return Result
}
