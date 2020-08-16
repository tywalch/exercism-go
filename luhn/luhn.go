package luhn

import (
	"strconv"
	"strings"
)

// Valid answers whether or not a give credit card number is valid
func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	if len(id) < 2 {
		return false
	}
	total := 0
	offset := len(id)%2 == 0
	for _, digit := range id {
		if value, err := strconv.Atoi(string(digit)); err == nil {
			if offset {
				value = value * 2
			}
			if value > 9 {
				value = value - 9
			}
			offset = !offset
			total += value
		} else {
			return false
		}
	}
	return total%10 == 0
}