package goreloaded

func Punc(s []string) []string {
	var result []string
	i := 0
	//grouping the punc
	for i < len(s) {
		if Only(s[i]) {
			group := s[i]
			j := i + 1
			for j < len(s) && Only(s[j]) {
				group += s[j]
				j++
			}
			result = append(result, group)
			i = j
		} else {
			result = append(result, s[i])
			i++
		}
	}
	//merge the puch with the previos word
	for j := 0; j < len(result); j++ {
		if Only(result[j]) {
			if j != 0 {
				result[j-1] = result[j-1] + result[j]
				result[j] = ""
			}
		}
	}
	var Result []string
	for _, r := range result {
		if r != "" {
			Result = append(Result, string(r))
		}
	}
		return Result
}
