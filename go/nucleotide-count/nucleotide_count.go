// Package dna counts nucleotides
package dna

import (
	"errors"
	"strings"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[DNA]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

const dna string = "CGAT"

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {

	var h Histogram
	for _, nucleotides := range d {
		if strings.Contains(dna, string(nucleotides)) {
			h[DNA(nucleotides)] += 1
		} else {
			return nil, errors.New("not a nucleontides of DNA")
		}

	}

	return h, nil
}
