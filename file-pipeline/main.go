// CodeCrafters — Operation Gopher Protocol
// Module: File Pipeline
// Author: [Faith Ejembi]
// Squad:  [Benchmarkers]

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go input.txt output.txt")
		return
	}
	faithInputName := os.Args[1]
	faithOutputName := os.Args[2]
	faithInputFile, myErr := os.Open(faithInputName)
	if myErr != nil {
		fmt.Println("Error opening input:", myErr)
		return
	}
	defer faithInputFile.Close()

	faithScanner := bufio.NewScanner(faithInputFile)
	var myLines []string
	textReadCount := 0
	textRemovedCount := 0

	for faithScanner.Scan() {
		myLine := faithScanner.Text()
		textReadCount++
		myLine = strings.TrimSpace(myLine)
		if myLine == "" || strings.Trim(myLine, "-") == "" {
			textRemovedCount++
			continue
		}
		if strings.Contains(myLine, "TODO") && myLine == strings.ToUpper(myLine) {
			myLine = strings.ReplaceAll(myLine, "TODO", "Action")
			myLine = strings.ToLower(myLine)
			myLine = strings.ToUpper(string(myLine[0])) + strings.ToLower(string(myLine[1:]))
		}
		if strings.Contains(myLine, "TODO") {
			myLine = strings.ReplaceAll(myLine, "TODO", "Action")
		}
		if myLine == strings.ToLower(myLine) {
			myLine = strings.ToUpper(myLine)
		}
		if myLine == strings.ToLower(myLine) {
			myLine = strings.ToUpper(string(myLine[0])) + strings.ToLower(string(myLine[1:]))
		}

		myLines = append(myLines, myLine)
	}

	outFile, myErr := os.Create(faithOutputName)
	if myErr != nil {
		fmt.Println("Error creating output:", myErr)
		return
	}
	defer outFile.Close()
	for index, line := range myLines {
		fmt.Fprintf(outFile, "%03d. %s\n", index+1, line)
	}

	fmt.Println("\n--- SUMMARY ---")
	fmt.Printf("Lines read: %d\n", textReadCount)
	fmt.Printf("Lines written: %d\n", len(myLines))
	fmt.Printf("Lines removed: %d\n", textRemovedCount)
	fmt.Println("RULES APPLIED:\n trimspace \n TODO - Action \n remove empty or blank space \n uppercase line to titlecase \n lowercase to uppercase")
	fmt.Println("Processing complete! Check:", faithOutputName)
}
