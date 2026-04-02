// CodeCrafters — Operation Gopher Protocol
// Module: String Transformer
// Author: [Your Name]
// Squad:  [Your Squad Name]

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(cap("wicked mubarak"))
	fmt.Println(title("the fall of the western power grid"))
	fmt.Println(snake("Operation Gopher Protocol"  ))
}

func upper(s string) string {
	if len(s) < 1 {
		return s

	}
	return strings.ToUpper(s)
}

func lower(s string) string {
	if len(s) < 1 {
		return s
	}
	return strings.ToLower(s)
}
func cap(s string) string {
	word := strings.Fields(s)
	for i, char := range word {
		word[i] = strings.ToUpper(string(char[0])) + strings.ToLower(string(char[1:]))

	}
	return strings.Join(word, " ")

}
func title(s string) string {
	word := strings.Fields(s)
	text := map[string]bool{
		"a":   true,
		"an":  true,
		"the": true,
		"and": true,
		"but": true,
		"or":  true,
		"for": true,
		"nor": true,
		"on":  true,
		"at":  true,
		"to":  true,
		"by":  true,
		"in":  true,
		"of":  true,
		"up":  true,
		"as":  true,
		"is":  true,
		"it":  true,
	}
	for i, item := range word {
		faithtext := strings.ToLower(item)
		if i == 0 || !text[faithtext] {
			word[i] = strings.ToUpper(string(faithtext[0])) + faithtext[1:]
		} else {
			word[i] = faithtext
		}

	}
	return strings.Join(word, " ")
}

func snake(s string) string {
	var ehi strings.Builder
	for _, char := range strings.ToLower(s) {
		switch {
		case char == ' ':
			ehi.WriteRune('_')
		case unicode.IsLetter(char) || unicode.IsDigit(char) || char == '_':
			ehi.WriteRune(char)
		}
	}
return ehi.String()
}
