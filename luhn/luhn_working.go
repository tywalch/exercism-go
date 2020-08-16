package luhn

// import (
// 	"errors"
// 	"strconv"
// 	"unicode"
// )

// func prepareDigits(id string) (string, error) {
// 	result := []rune{}
// 	for _, num := range id {
// 		if unicode.IsDigit(num) {
// 			result = append(result, num)
// 		} else if unicode.IsSpace(num) {
// 			continue
// 		} else {
// 			return "", errors.New("Invalid character")
// 		}
// 	}
// 	return string(result), nil
// }

// // Valid answers whether or not a give credit card number is valid
// func Valid(id string) bool {
// 	value, err := prepareDigits(id)
// 	if err != nil || len(value) < 2 {
// 		return false
// 	}
// 	offset := len(value)%2 != 0
// 	total := 0
// 	for i, num := range value {
// 		if digit, err := strconv.Atoi(string(num)); err == nil {
// 			// Offset index if doubling
// 			if offset {
// 				i++
// 			}

// 			// double every second digit
// 			if i%2 == 0 {
// 				digit = digit * 2
// 			}

// 			// If doubling the number results in a number greater than 9 then subtract 9 from the product.
// 			if digit > 9 {
// 				digit = digit - 9
// 			}

// 			total += digit
// 		}
// 	}
// 	return total%10 == 0
// }
