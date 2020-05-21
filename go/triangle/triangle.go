// Package trianle detects which trianle is
package triangle

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

	return k
}
