// Package raindrops converts digits to raindrops sound
package raindrops

import "strconv"

// Convert converts a digit to a raindrop sound
func Convert(num int) string {
	// raindropsSound a string of raindrops
	var raindropsSound string
	if num%3 == 0 {
		raindropsSound += "Pling"
	}
	if num%5 == 0 {
		raindropsSound += "Plang"
	}
	if num%7 == 0 {
		raindropsSound += "Plong"
	}
	if raindropsSound == "" {
		raindropsSound = strconv.Itoa(num)
	}
	return raindropsSound
}
