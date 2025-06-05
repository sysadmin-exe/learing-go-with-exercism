package thefarm

import "errors"
import "fmt"

// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, cowNumber int) (float64, error) {
    fAmount, err := fc.FodderAmount(cowNumber)
    if err != nil {
        return 0, err
    }
    fFactor, err := fc.FatteningFactor()
    if err != nil {
        return 0, err
    }
    fpc := (fAmount / float64(cowNumber)) * fFactor
    return fpc, nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, cowNumber int) (float64, error) {
    if cowNumber > 0 {
        return DivideFood(fc, cowNumber)
    } else {
        return 0, errors.New("invalid number of cows")
    }
}

type InvalidCowsError struct {
  cow int
  details string
}

func (e *InvalidCowsError) Error() string {
  return fmt.Sprintf("%d cows are invalid: %s", e.cow, e.details)
}


// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(cowNumber int) error {
    if cowNumber < 0 {
        return &InvalidCowsError{
            cow: cowNumber,
            details: "there are no negative cows",
        }
    } else if cowNumber == 0 {
        return &InvalidCowsError{
            cow: cowNumber,
            details: "no cows don't need food",
        }
    } else {
        return nil
    }
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
