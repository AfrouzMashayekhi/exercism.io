// Package hamming works on hamming distance algorithm
package hamming

import (
	"errors"
)

// Distance calculate hamming distance return hamming distance or error
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("can't compare two string")
	}
	// hammingDistance counts the difference of two string
	hammingDistance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			hammingDistance++
		}
	}
	return hammingDistance, nil

}
