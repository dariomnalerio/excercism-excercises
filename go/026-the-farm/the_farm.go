package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(fc FodderCalculator, cows int) (float64, error) {
	totalFodderPerCow, err := fc.FodderAmount(1)
	if err != nil {
		return 0.0, err
	}
	totalFodder := totalFodderPerCow

	fatteningFactor, err := fc.FatteningFactor()
	if err != nil {
		return 0.0, err
	}

	return (totalFodder * fatteningFactor) / float64(cows), nil
}

func ValidateInputAndDivideFood(fc FodderCalculator, cows int) (float64, error) {
	var ErrInvalidInt = errors.New("invalid number of cows")
	if cows > 0 {
		dividedFood, err := DivideFood(fc, cows)
		if err != nil {
			return 0, err
		}
		return dividedFood, nil
	}
	return 0, ErrInvalidInt
}

type InvalidCowsError struct {
	cows    int
	message string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.message)
}

func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{
			cows:    cows,
			message: "there are no negative cows",
		}
	} else if cows > 0 {
		return nil
	}
	return &InvalidCowsError{
		cows:    cows,
		message: "no cows don't need food",
	}
}
