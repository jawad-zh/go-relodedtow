package goreloaded

func AtoAn(s []string) []string {
	for i := 0; i < len(s); i++ {
		if (Isa(s[i]) && i != len(s)-1 && Vowel(string(s[i+1][0])) )  {
			s[i] = "An"
		}
	}
	return s
}
