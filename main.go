package main

import (
	"fmt"
	"os"
	"strings"

	goreloaded "goreloaded/functionss"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <input-file> <output-file>")
		return
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]
	if strings.HasSuffix(strings.ToLower(outputFile), ".go") {
		fmt.Println("Error: Output file cannot have a '.go' extension!")
		return
	}
	if inputFile == outputFile {
		fmt.Println("Error: Input and output files cannot be the same!")
		return
	}

	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}
	txt := string(content)

	clean := goreloaded.Clean(txt)
	// clean = goreloaded.MergeFlags(clean)
	clean = goreloaded.Convert(clean)
	txt = strings.Join(clean, " ")
	clean = goreloaded.SplitPunc(txt)
	clean = goreloaded.Punc(clean)
	clean = goreloaded.AtoAn(clean)
	txt = strings.Join(clean, " ")
	clean = goreloaded.Quote(txt)
	finalOutput := goreloaded.JoinCleaned(clean)
	// fmt.Printf("%q",finalOutput)
	err = os.WriteFile(outputFile, []byte(finalOutput), 0644)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
		return
	}
	fmt.Printf("Processing complete. Output written to %s\n", outputFile)
}
