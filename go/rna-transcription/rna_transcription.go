//Package strand transform DNA to RNA
package strand

// ToRNA get a string of dna and output rna string
func ToRNA(dna string) string {
	//rna output rna string
	var rna string
	// mapping of dna to rna
	RNAmapping := map[string]string{
		`G`: `C`,
		`C`: `G`,
		`T`: `A`,
		`A`: `U`,
	}

	for _, nucleotides := range dna {
		rna += RNAmapping[string(nucleotides)]

	}
	return rna
}
