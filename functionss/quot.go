package goreloaded

import (
	"strings"
)

func Quote(tokens []string) []string {
	var result []string
	i := 0
	for i < len(tokens) {
		if tokens[i] == "'" {
			// Find closing quote or end
			j := i + 1
			for j < len(tokens) && tokens[j] != "'" {
				j++
			}

			if j == len(tokens) {
				// No closing quote found - take till end
				j = len(tokens)
			}

			insideTokens := tokens[i+1 : j]

			// Case 1: if insideTokens are all single quotes or empty, count pairs of quotes
			if allSingleQuotes(insideTokens) {
				count := len(insideTokens)
				for k := 0; k < count/2; k++ {
					result = append(result, "''")
				}
				if count%2 == 1 {
					result = append(result, "'")
				}
			} else {
				// Trim spaces inside each token before joining
				for k := range insideTokens {
					insideTokens[k] = strings.TrimSpace(insideTokens[k])
				}
				joined := strings.Join(insideTokens, " ")
				joined = strings.TrimSpace(joined)
				quoted := "'" + joined + "'"
				result = append(result, quoted)
			}

			i = j + 1
		} else {
			result = append(result, tokens[i])
			i++
		}
	}
	return result
}

func allSingleQuotes(tokens []string) bool {
	for _, t := range tokens {
		if t != "'" {
			return false
		}
	}
	return true
}
