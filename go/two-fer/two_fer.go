// Package twofer implements shares owner
package twofer

import "fmt"

// ShareWith returns who get one.
func ShareWith(name string) string {

	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
