// Package proverb convert a rhyme to a proverb
package proverb

import "fmt"

// Proverb get a rhyme transforms it to a proverb
func Proverb(rhyme []string) []string {
	var proverb []string
	if len(rhyme) == 0 {
		return nil
	}
	//iterate over rhyme
	for i := 0; i < len(rhyme)-1; i++ {
		proverb = append(proverb, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))

	}
	// the last line of proverb
	proverb = append(proverb, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))

	return proverb
}
