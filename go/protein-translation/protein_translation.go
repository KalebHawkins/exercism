package protein

import "fmt"

var translation = map[string][]string{
	"Methionine":    {"AUG"},
	"Phenylalanine": {"UUU", "UUC"},
	"Leucine":       {"UUA", "UUG"},
	"Serine":        {"UCU", "UCC", "UCA", "UCG"},
	"Tyrosine":      {"UAU", "UAC"},
	"Cysteine":      {"UGU", "UGC"},
	"Tryptophan":    {"UGG"},
	"STOP":          {"UAA", "UAG", "UGA"},
}

var ErrStop = fmt.Errorf("terminating codon recieved")
var ErrInvalidBase = fmt.Errorf("invalid protien recieved")

func FromRNA(rna string) ([]string, error) {
	var protiens = make([]string, 0)

	for i := 0; i < len(rna); i += 3 {
		protien, err := FromCodon(rna[i : i+3])

		if err == ErrStop {
			return protiens, nil
		}
		if err == ErrInvalidBase {
			return protiens, err
		}

		protiens = append(protiens, protien)
	}

	return protiens, nil
}

func FromCodon(codon string) (string, error) {
	for k, v := range translation {
		for _, c := range v {
			if c == codon && k == "STOP" {
				return "", ErrStop
			}
			if c == codon {
				return k, nil
			}
		}
	}

	return "", ErrInvalidBase
}
