// Package accumulate operates some functions on strings
package accumulate

// Accumulate converts str to a string depend on input operator
func Accumulate(str []string, converter func(string) string) []string {
	// Accumulated returns array of string that operator applied
	var accumulated []string
	for _, words := range str {
		accumulated = append(accumulated, converter(words))
	}
	return accumulated
}
