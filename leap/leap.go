
package leap

// IsLeapYear reports if year provided is a leap year.
func IsLeapYear(year int) bool {
	if year % 100 == 0 {
		return year % 400 == 0
	} else {
		return year % 4 == 0
	}
}
