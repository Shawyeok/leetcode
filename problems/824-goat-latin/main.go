package main

import (
	"strings"
)

func main() {
}

func toGoatLatin(S string) string {
	vowels := map[byte]int{
		'a': 1,
		'e': 1,
		'i': 1,
		'o': 1,
		'u': 1,
		'A': 1,
		'E': 1,
		'I': 1,
		'O': 1,
		'U': 1,
	}
	words := strings.Split(S, " ")
	goatWords := []string{}
	for i, word := range words {
		first := word[0]
		var goatWord string
		if vowels[first] == 1 {
			goatWord = word + "ma" + strings.Repeat("a", i+1)
		} else {
			goatWord = word[1:] + word[:1] + "ma" + strings.Repeat("a", i+1)
		}

		goatWords = append(goatWords, goatWord)
	}

	return strings.Join(goatWords, " ")
}
