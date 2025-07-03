package goreloaded

func StringFlag(s string)bool{
 for i := 0; i < len(s); i++ {
		if ((s) == "(up," && i+1 < len(s))  {
			return true
		}else if ((s) == "(low," && i+1 < len(s))  {
			return true
		}else if ((s) == "(cap," && i+1 < len(s))  {
			return true
		}else if ((s) == "(hex," && i+1 < len(s)) {
			return true
		}else if ((s) == "(bin," && i+1 < len(s)) {
			return true
		}
	}
	return false
}