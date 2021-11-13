package romannumerals

import (
	"errors"
)

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input > 3000 {
		return "", errors.New("input out of range")
	}

	values := []int{
		1000, 900, 500, 400,
		100, 90, 50, 40,
		10, 9, 5, 4, 1,
	}

	symbols := []string{
		"M", "CM", "D", "CD",
		"C", "XC", "L", "XL",
		"X", "IX", "V", "IV",
		"I"}

	var roman string
	for i := 0; input > 0; i++ {

		val := input / values[i]

		for j := 0; j < val; j++ {
			roman += symbols[i]
			input -= values[i]
		}
	}
	return roman, nil
}
