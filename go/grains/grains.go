package grains

import (
	"errors"
	"math"
)

// const uintMax uint64 = 18446744073709551615

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("number not within range: 1-64")
	}

	return uint64(1 << uint(number-1)), nil
}

func Total() uint64 {
	return math.MaxUint64
}
