package romannumerals

import (
	"errors"
)

type romanNumeral struct {
	arabic int
	roman  string
}

var romanNumerals = []romanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// ToRomanNumeral returns the roman numeral representation of positive whole numbers less than or equal to 3000
func ToRomanNumeral(i int) (string, error) {
	if i <= 0 || i > 3000 {
		return "", errors.New("provided integers must be positive whole numbers less than or equal to 3000")
	}
	roman := ""
	for _, num := range romanNumerals {
		for ; num.arabic <= i; i -= num.arabic {
			roman += num.roman
		}
	}
	return roman, nil
}