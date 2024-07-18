// package leap is an exercise for determining leap years.
package leap

// IsLeapYear returns true if the year is a leap year, false otherwise.
func IsLeapYear(year int) bool {
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		return true
	}
	return false
}
