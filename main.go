package main

import (
	"fmt"
	"os"
	"strings"

	goreloaded "goreloaded/functionss"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <input-file.txt> <output-file.txt>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]
	if !strings.HasSuffix(strings.ToLower(inputFile), ".txt") {
		fmt.Println("Error: You can only run this program with a .txt input file!")
		return
	}
	if strings.HasSuffix(strings.ToLower(outputFile), ".go") {
		fmt.Println("Error: Output file cannot have a '.go' extension!")
		return
	}
	if !strings.HasSuffix(strings.ToLower(outputFile), ".txt") {
		fmt.Println("Error: Output file must have a '.txt' extension!")
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
	contentSplit := strings.Split(txt, "\n")
	transfor := ""
	for i, x := range contentSplit {
		next := goreloaded.Process(x)
		transfor += next
		if i != len(contentSplit) {
			transfor += " \n"
		}
	}
	err = os.WriteFile(outputFile, []byte(transfor), 0o644)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
		return
	}

	fmt.Printf("Processing complete. Output written to %s\n", outputFile)
}
