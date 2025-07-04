package goreloaded



func AtoAn(s []string) []string {
	for i := 0; i < len(s); i++ {
		if (Isa(s[i]) && i != len(s)-1 && Vowel(string(s[i+1][0])) )  {
			if s[i] == "a" {
				s[i] = "an"
			}else if s[i] == "A"{
				s[i] = "An"
			}
		}
	}
	return s
}
