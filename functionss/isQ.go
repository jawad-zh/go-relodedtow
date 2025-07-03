package goreloaded

func IsQ(s string)bool{
	hasQuote:=false
	for i:=0 ; i < len(s) ; i++{
		if s[i] == '\''{
			hasQuote = true
			break
		}else{
			hasQuote = false
		}
	}
	return hasQuote
}