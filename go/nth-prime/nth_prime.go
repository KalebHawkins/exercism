package prime

import (
	"fmt"
	"math"
)

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n <= 0 {
		return 0, fmt.Errorf("integer should be greater than 0")
	}

	primePlace := 0
	primeNum := 2
	for i := primeNum; ; i++ {
		// Is the number greater than 2 and divisible by 2?
		// If so move on.
		if i > 2 && i%2 == 0 {
			continue
		}

		// Is the number greater than 3 and divisible by 3?
		// If so move on.
		if i%3 == 0 && i > 3 {
			continue
		}

		// Is the number divisible by 5 and greater than 5?
		// If so move on.
		if (i%5 == 0) && i > 5 {
			continue
		}

		startPoint := int(math.Floor(math.Sqrt(float64(i))))
		var isPrime bool = true
		for j := startPoint; j > 1; j-- {
			if i%j == 0 {
				isPrime = false
				break
			}
		}

		if !isPrime {
			continue
		}

		primePlace++
		primeNum = i

		if primePlace == n {
			return primeNum, nil
		}
	}
}
