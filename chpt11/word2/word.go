// Package word provides utilities for word games.
package word

import (
	"math/rand"
	"strings"
	"unicode"
)

// IsPalindrome reports whether s reads the same forward and backward.
// Letter case is ignored, as are non-letters.
func IsPalindrome(s string) bool {
	strings.Join
	var letters = make([]rune, 0, len(s))
	for _, l := range s {
		if unicode.IsLetter(l) {
			letters = append(letters, unicode.ToLower(l))
		}
	}
	n := len(letters) / 2
	for i := 0; i < n; i++ {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}

// randomPalindrome returns a palindrome whose length and contents
// are derived from the pseudo-random number generator rng.
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		if unicode.IsLetter(r) {
			runes[i] = unicode.ToLower(r)
			runes[n-1-i] = unicode.ToLower(r)
		}
	}
	return string(runes)
}

// randomNonPalindrome returns a string whose length and contents
// are derived from the pseudo-random number generator rng.
func randomNonPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n+2)
	for i := 0; i < n+2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999
		if unicode.IsLetter(r) {
			runes[i] = unicode.ToLower(r)
			runes[n+2-1-i] = unicode.ToLower(r / rune(rng.Intn(10)+4))
		}
	}
	return string(runes)
}
