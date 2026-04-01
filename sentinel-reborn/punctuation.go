package main

import (
	"fmt"
	"strings"
)

func fixPunctuations(text string) string {
	text = strings.ReplaceAll(text, " .", ".")
	text = strings.ReplaceAll(text, " ?", "?")
	text = strings.ReplaceAll(text, ". ..", "... ")
	text = strings.ReplaceAll(text, " ;", ";")
	text = strings.ReplaceAll(text, " :", ":")
	text = strings.ReplaceAll(text, " ,", ", ")
	text = strings.ReplaceAll(text, " !", "!")
	return text
}

func main() {
	fmt.Println("I was sitting over there ,and then BAMM !!")
	fmt.Println(fixPunctuations("omo director x .... my brain is not braining anymore!!"))
	fmt.Println(fixPunctuations("bench markers is the best ... squad ever no doubt"))
	fmt.Println(fixPunctuations("I was sitting over there ,and then BAMM !!"))
}
