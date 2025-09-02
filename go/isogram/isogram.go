package isogram

import "unicode"

// IsIsogram checks if a word is an isogram (a word with no repeating letters).
// It ignores letter case, hyphens, and spaces.
func IsIsogram(word string) bool {
	letters := make(map[rune]bool)
	for _, char := range word {
		if char == '-' || char == ' ' { // allow hyphens and spaces
			continue
		}
		char = unicode.ToLower(char) // handle case insensitivity

		if letters[char] {
			return false // found a duplicate letter
		}
		letters[char] = true
	}
	return true
}
