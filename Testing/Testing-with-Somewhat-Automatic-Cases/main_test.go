package main

import "testing"

var TestCases = []struct {
	TestName string
	Dividend float64
	Divider  float64
	Expected float64
	IsErr    bool
}{
	{"valid", 10.0, 1.0, 10.0, false},
	{"valid", 10.0, 2.0, 5.0, false},
	{"invalid", 10.0, 0, 0, true},
}

func TestDivideBy(t *testing.T) {
	// Now this is for the semiautomatic testing

	for _, tt := range TestCases {
		// This is a for loop to check every case of test cases

		result, err := divideBy(tt.Dividend, tt.Divider)
		// This runs the function divideBy using every dividend and divider in the test case

		if tt.IsErr {
			if err == nil {
				t.Error("Expected error but did not get one")
			}
		} else {
			if err != nil {
				t.Error("Did not expect error but got one")
			}
		}
		// These if else decision is to check whether the error thrown by the function is the same as expected error

		if result != tt.Expected {
			t.Error("Wrong result")
		}
		// This tests whether the result of the function is the same as the test cases
	}
}
