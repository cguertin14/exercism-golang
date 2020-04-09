// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

func DoesSentenceContainLetters(sentence string) bool {
	for _, letter := range sentence {
		if unicode.IsLetter(letter) {
			return true
		}
	}

	return false
}

func IsSentenceInAllCaps(sentence string) bool {
	for _, letter := range sentence {
		if !unicode.IsUpper(letter) && unicode.IsLetter(letter) {
			return false
		}
	}

	return true
}

func IsEmpty(sentence string) (string, bool) {
	trimmed := strings.TrimSpace(sentence)
	return trimmed, trimmed == ""
}

func IsSentenceAQuestion(sentence string) bool {
	trimmed, IsEmpty := IsEmpty(sentence)
	if IsEmpty {
		return false
	}

	return string(trimmed[len(trimmed)-1]) == "?"
}

func Hey(remark string) string {
	var toReturn string
	IsAllCaps := IsSentenceInAllCaps(remark)
	IsAQuestion := IsSentenceAQuestion(remark)
	ContainsLetters := DoesSentenceContainLetters(remark)
	_, IsEmpty := IsEmpty(remark)

	if IsAllCaps && !IsAQuestion && ContainsLetters {
		toReturn = "Whoa, chill out!"
	} else if IsAQuestion && IsAllCaps && ContainsLetters {
		toReturn = "Calm down, I know what I'm doing!"
	} else if IsAQuestion {
		toReturn = "Sure."
	} else if IsEmpty {
		toReturn = "Fine. Be that way!"
	} else {
		toReturn = "Whatever."
	}

	return toReturn
}
