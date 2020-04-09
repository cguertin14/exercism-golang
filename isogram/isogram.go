package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	if word == "" {
		return true
	}

	exists := map[rune]bool{}
	word = strings.ToLower(word)

	for _, char := range word {
		if unicode.IsLetter(char) && exists[char] {
			return false
		}

		exists[char] = true
	}

	return true
}
