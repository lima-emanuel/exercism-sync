package thefarm

import (
	"errors"
	"fmt"
)

type InvalidCowsError struct {
	nc      int
	message string
}

func DivideFood(fc FodderCalculator, nc int) (float64, error) {
	total_fodder, err := fc.FodderAmount(nc)
	if err != nil {
		return 0, err
	}

	fat_factor, err := fc.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return total_fodder * fat_factor / float64(nc), nil
}

func ValidateInputAndDivideFood(fc FodderCalculator, nc int) (float64, error) {
	if nc > 0 {
		return DivideFood(fc, nc)
	} else {
		return 0.0, errors.New("invalid number of cows")
	}
}

func ValidateNumberOfCows(nc int) error {
	switch {
	case nc < 0:
		return &InvalidCowsError{nc, "there are no negative cows"}
	case nc == 0:
		return &InvalidCowsError{nc, "no cows don't need food"}
	default:
		return nil
	}
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.nc, e.message)
}
