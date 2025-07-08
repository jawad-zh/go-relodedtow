package goreloaded

import "strings"

func Prossec(txt string) string {
	clean := Clean(txt)
	clean = Convert(clean)
	txt = strings.Join(clean," ")
	clean = SplitPunc(txt)
	clean = Punc(clean)
	txt = strings.Join(clean, " ")
	clean = strings.Fields(Quotes(txt))
	clean = AtoAn(clean)
	txt = strings.Join(clean, " ")
	return txt
}
