package hamming

import (
	"fmt"
)

func Distance(a, b string) (int, error) {
	ar := []rune(a)
	br := []rune(b)

	if len(ar) != len(br) {
		return 0, fmt.Errorf("invalid sequences: sequences cannot be of different length")
	}

	var dist int
	for i := range ar {
		if ar[i] != br[i] {
			dist++
		}
	}

	return dist, nil
}
