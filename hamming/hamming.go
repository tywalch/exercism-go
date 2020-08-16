package hamming

import "errors"

// Distance takes two strings of equal length and returns the Hamming Distance between them
func Distance(a, b string) (int, error) {
	dist := 0
	aRunes := []rune(a)
	bRunes := []rune(b)
	if len(aRunes) != len(bRunes) {
		return dist, errors.New("provided values must have the same length")
	}
	for i := range aRunes {
		if aRunes[i] != bRunes[i] {
			dist++
		}
	}
	return dist, nil
}
