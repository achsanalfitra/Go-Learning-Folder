package main

import (
	"errors"
	"fmt"
)

const dividend float64 = 10.0
const divider float64 = 0

func main() {
	divided, err := divideBy(dividend, divider)
	if err != nil {
		fmt.Println(err)
		return
	}
	// This is the method of testing when the function is run in the main program
	// The function will have its own error catching where the error output can be caught
	// When the error is caught, the entire main function won't continue

	fmt.Println(divided)
}

func divideBy(x, y float64) (float64, error) {
	// This is the structure of a function that can be tested for its errors
	// It gives the error as an output
	// If when the function has been run it doesn't return any error or nil, then it works

	var result float64

	if y == 0 {
		return result, errors.New("You cannot divide by zero")
	}
	// This is the part where if the error case is satisfied, when the divider is equal to zero, an error is returned
	// The error is created by calling the errors package with the function New

	result = x / y
	return result, nil
	// Otherwise, when there were no errors, this function returns nil as its error
}
