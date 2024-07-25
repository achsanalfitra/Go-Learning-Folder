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

	fmt.Println(divided)
}

func divideBy(x, y float64) (float64, error) {
	var result float64

	if y == 0 {
		return result, errors.New("You cannot divide by zero")
	}

	result = x / y
	return result, nil
}

// This is the same code used with the main.go file in the Catching an Error in main
