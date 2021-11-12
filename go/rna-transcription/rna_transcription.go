package strand

func ToRNA(dna string) string {
	var rnaStrand string

	for _, v := range dna {
		switch v {
		case 'G':
			rnaStrand += "C"
		case 'C':
			rnaStrand += "G"
		case 'T':
			rnaStrand += "A"
		case 'A':
			rnaStrand += "U"
		}
	}

	return rnaStrand
}
