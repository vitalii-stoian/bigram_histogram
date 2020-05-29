package main

import "strings"

// Parse takes a text 's' and splits it into words.
// Any non-ASCII letters a-z A-Z and digits 0-9 are ignored and treated as words delimiters.
//
// Complexity: O(N) where N is a number of characters in the input text.
//
func Parse(s string) []string {
	words := []string{}
	wordStartPos := 0
	i := 0

	// Only allow ASCII and digits.
	isLetter := func(c uint8) bool {
		isLowercase := c >= 'a' && c <= 'z'
		isUppercase := c >= 'A' && c <= 'Z'
		isDigit := c >= '0' && c <= '9'
		return isLowercase || isUppercase || isDigit
	}

	// Skip whitespaces in the front before the first word.
	for i < len(s) && !isLetter(s[i]) {
		i++
	}

	wordStartPos = i
	for i < len(s) {
		// Skip any non-whitespace characters in the current word.
		if isLetter(s[i]) {
			i++
			continue
		}

		// Add a new word (converted to lowercase). Its starting position is tracked by the wordStartPos.
		words = append(words, strings.ToLower(s[wordStartPos:i]))
		i++

		// Skip whitespaces between words.
		for i < len(s) && !isLetter(s[i]) {
			i++
		}

		// Memorize the starting position of the next word.
		wordStartPos = i
	}

	// Last word could end at the end of the input text if no whitespace characters are at the end.
	if wordStartPos < len(s) {
		words = append(words, strings.ToLower(s[wordStartPos:]))
	}
	return words
}
