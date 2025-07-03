package goreloaded

func Ingroup(s string) bool{
	for _,char:=range s {
		if char != '.' && char != ',' && char != '!' && char != '?' && char != ':' && char != ';'  {
			return false
			break
		}
	}
	return true
}