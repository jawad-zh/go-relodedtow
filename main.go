package main

import (
	"fmt"
	goreloaded "goreloaded/functionss"
	"os"
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
	clean = goreloaded.Quote(clean)
	fmt.Print(goreloaded.JoinCleaned(clean))	
	// fmt.Printf("%q",clean)
	// fmt.Print(goreloaded.Tst(txt))
	// fmt.Print(len(clean))
	// fmt.Print(clean)
}
