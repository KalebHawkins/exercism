package hamming

import (
	"errors"
)

// Distance returns the count of differences between to strings of the same length.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("sequences have different lengths")
	}

	count := 0
	for i, v := range a {
		if byte(v) != b[i] {
			count++
		}
	}

	return count, nil

}
