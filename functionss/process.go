package goreloaded

import "strings"

func Prossec(txt string) string {
	clean := Clean(txt)
	clean = Convert(clean)
	txt = JoinCleaned(clean)
	clean = SplitPunc(txt)
	clean = Punc(clean)
	txt = strings.Join(clean, " ")
	clean = strings.Fields(QuoteFmt(txt))
	clean = AtoAn(clean)
	txt = JoinCleaned(clean)
	return txt
}
