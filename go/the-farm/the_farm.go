package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	message string
}

func (s *SillyNephewError) Error() string {
	return s.message
}

func NewSillyNephewError(cows int) *SillyNephewError {
	return &SillyNephewError{message: fmt.Sprintf("silly nephew, there cannot be %d cows", cows)}
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	f, err := weightFodder.FodderAmount()

	switch {
	case err == ErrScaleMalfunction && f > 0:
		return f * 2 / float64(cows), nil
	case err == ErrScaleMalfunction && f < 0:
		return 0, errors.New("negative fodder")
	case err != nil:
		return 0, errors.New("non-scale error")
	case f < 0:
		return 0, errors.New("negative fodder")
	case cows == 0:
		return 0, errors.New("division by zero")
	case cows < 0:
		return 0, NewSillyNephewError(cows)
	}

	return f / float64(cows), nil
}
