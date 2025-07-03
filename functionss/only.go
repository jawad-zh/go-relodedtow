package goreloaded

func Only(s string)bool{
	for _,char:=range s {
		if char != '.' && char != ',' && char != '!' && char != '?' && char != ':' && char != ';'  {
			return false
		}
	}
	return true
}