package main

import (
	"fmt"
	"os"
	"strings"

	goreloaded "goreloaded/functionss"
)

func main() {
	content, err := os.ReadFile("input/sampl.txt")
	if err != nil {
		fmt.Print("Error:", err)
		return
	}
	txt := string(content)
	clean := goreloaded.Clean(txt)
	clean = goreloaded.MergeFlags(clean)
	clean = goreloaded.Convert(clean)
	clean = goreloaded.Punc(clean)
	clean = goreloaded.AtoAn(clean)
	txt = strings.Join(clean, " ")
	clean = goreloaded.Quote(txt)
	fmt.Print(goreloaded.JoinCleaned(clean))
	// fmt.Printf("%q",clean)
	// fmt.Print(txt)
	// fmt.Print(goreloaded.Tst(txt))
	// fmt.Print(len(clean))
	// fmt.Print(clean)
}
