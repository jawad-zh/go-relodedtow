package goreloaded

func StringFlag(s1, s2 string) bool {

	if s1 == "(up," && CorrectEnd(s2) {
		return true
	} else if s1 == "(low," && CorrectEnd(s2) {
		return true
	} else if s1 == "(cap," && CorrectEnd(s2) {
		return true
	} else if s1 == "(hex," && CorrectEnd(s2) {
		return true
	} else if s1 == "(bin," && CorrectEnd(s2) {
		return true
	}

	return false
}
