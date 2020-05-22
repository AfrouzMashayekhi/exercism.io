// Package trianle detects which trianle is
package triangle

import "math"

// Kind a type for your triangle
type Kind string

// define names of triangle
const (
	NaT Kind = "Nat"
	Equ Kind = "Equ"
	Iso Kind = "Iso"
	Sca Kind = "Sca"
)

// KindFromSides detects triangle type and return its kind
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	if a+b+c <= 0 || math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) || math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		k = NaT
	} else {
		switch {
		case a == b && b == c:
			k = Equ
		case a > b+c || b > a+c || c > b+a:
			k = NaT
		case a == b || b == c || a == c:
			k = Iso
		default:
			k = Sca
		}
	}

	return k
}
