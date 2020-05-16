// leap package state a year is a leap year or no
package leap

// Divisable return if a num divisable true
func Divisable(num int, mod int) bool {

	return (num%mod == 0)
}

// IsLeapYear shows that a year is a loeap year or not
func IsLeapYear(year int) bool {
	return Divisable(year, 4) && (!Divisable(year, 100)) || Divisable(year, 400)
}
