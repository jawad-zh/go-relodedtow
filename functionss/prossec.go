package goreloaded

import (
	"strings"
)

func Process(txt string) string {
	clean := Clean(txt)
	clean = Convert(clean)
	txt = strings.Join(clean, " ")
	clean = SplitPunc(txt)
	clean = Punc(clean)
	clean = AtoAn(clean)
	txt = strings.Join(clean, " ")
	clean = Quote(txt)
	finalOutput := JoinCleaned(clean)
	return finalOutput
}
