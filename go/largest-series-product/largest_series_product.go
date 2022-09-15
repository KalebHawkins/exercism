package lsproduct

import (
	"fmt"
	"unicode"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span < 0 {
		return -1, fmt.Errorf("invalid: negative span")
	}
	if len(digits) < span {
		return -1, fmt.Errorf("invalid: string longer the span")
	}

	for _, v := range digits {
		if !unicode.IsNumber(v) {
			return 0, fmt.Errorf("invalid: %v is not a number", v)
		}
	}

	var largestProduct int
	for i := 0; i <= len(digits)-span; i++ {
		product := 1
		for j := 0; j < span; j++ {
			product *= int(digits[i+j] - '0')
		}

		if product > largestProduct {
			largestProduct = product
		}
	}

	return int64(largestProduct), nil
}
