package goutilstest

import (
	"strings"
	"unicode"
)

// Reverse reverses a string
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// ToTitleCase converts a string to title case
func ToTitleCase(s string) string {
	return strings.Title(s)
}

// IsPalindrome checks if a string is a palindrome
func IsPalindrome(s string) bool {
	clean := ""
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			clean += string(unicode.ToLower(r))
		}
	}
	return clean == Reverse(clean)
}
