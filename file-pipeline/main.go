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
		lowerLine := strings.ToLower(myLine)
		if strings.Contains(lowerLine, "todo") {
			myLine = strings.ReplaceAll(myLine, "TODO", "✦ ACTION")
			myLine = strings.ReplaceAll(myLine, "todo", "✦ ACTION")
		}
		if strings.Contains(lowerLine, "classified") {
			myLine = strings.ReplaceAll(myLine, "CLASSIFIED", "[REDACTED]")
			myLine = strings.ReplaceAll(myLine, "classified", "[REDACTED]")
		}

		myLines = append(myLines, myLine)
	}

	outFile, myErr := os.Create(faithOutputName)
	if myErr != nil {
		fmt.Println("Error creating output:", myErr)
		return
	}
	defer outFile.Close()

	     fmt.Fprintln(outFile, "SENTINEL REPORT - PROCESSED")
	     fmt.Fprintln(outFile, "")

	for index, line := range myLines {
		fmt.Fprintf(outFile, "%03d. %s\n", index+1, line)
	}

	     fmt.Fprintln(outFile, "\n--- SUMMARY ---")
	     fmt.Fprintf(outFile, "Lines read: %d\n", textReadCount)
	     fmt.Fprintf(outFile, "Lines written: %d\n", len(myLines))
	     fmt.Fprintf(outFile, "Lines removed: %d\n", textRemovedCount)

	fmt.Println("Processing complete! Check:", faithOutputName)
}
