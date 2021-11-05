package leap

// IsLeapYear returns true is a given year is a leap year or false if it is not.
func IsLeapYear(y int) bool {
	return y%4 == 0 && !(y%100 == 0) || y%400 == 0
}
