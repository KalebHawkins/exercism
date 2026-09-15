package hamming

import (
	"fmt"
	"strings"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("invalid sequences: sequences of different length")
	}

	upperA := strings.ToUpper(a)
	upperB := strings.ToUpper(b)

	var dist int
	for i, v := range upperA {
		if v != rune(upperB[i]) {
			dist++
		}
	}

	return dist, nil
}
