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

	if !strings.HasSuffix(outputFile, ".txt") {
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
	split := strings.Split(txt, "\n")
	NewData := ""
	for i, x := range split {
		NewData += goreloaded.Prossec(x)
		if i != len(split)-1 {
			NewData += "\n"
		}
	}
	err = os.WriteFile(outputFile, []byte(NewData), 0o644)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
		return
	}
	fmt.Printf("Processing complete. Output written to %s\n", outputFile)
}
