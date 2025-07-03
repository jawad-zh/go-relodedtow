package goreloaded

func MergeFlags(s []string) []string {
	for i:=0 ; i < len(s) ; i++{
		if (s[i] == "(up" || s[i] == "(low" || s[i] == "(cap") && i != len(s)-1 && s[i+1] == ","{
			s[i] = s[i] + s[i+1]
			s[i+1] = ""
		}
	}
	var Result []string
	for _,r:= range s{
		if r != ""{
			Result = append(Result, r)
		}
	}
	return Result
}
