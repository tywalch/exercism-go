package isogram

import "unicode"

// IsIsogram checks a passed string for reoccuring letters.
func IsIsogram(word string) bool {
	encountered := map[rune]struct{}{}
	for _, letter := range word {
		if unicode.IsLetter(letter) {
			letter = unicode.ToLower(letter)	
			if _, exists := encountered[letter]; exists {
				return false
			} 
			encountered[letter] = struct{}{}
		}
	}
	return true
}