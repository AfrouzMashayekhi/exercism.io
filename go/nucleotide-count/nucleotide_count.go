// Package dna counts nucleotides
package dna

import (
	"errors"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	// rune can not be empty
	h := Histogram{
		'C': 0,
		'G': 0,
		'T': 0,
		'A': 0,
	}
	for _, nucleotides := range d {
		if _, ok := h[nucleotides]; ok {
			h[nucleotides] += 1
		} else {
			return nil, errors.New("not a nucleontides of DNA")
		}

	}

	return h, nil
}
