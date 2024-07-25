package main

import "testing"

func TestDivideBy(t *testing.T) {
	// This is a standard function with test
	// This file runs by using command go test
	// The command would only work if there was a module file in the parent folder of this
	// THe name of this file "main_test.go" follows the rule of main is the file that we want to test and _test.go is the suffix for testing file
	// The function name also follows similar convention by adding the name Test as a prefix
	_, err := divideBy(10.0, 1.0)
	if err != nil {
		t.Error("This violates the rule")
	}
}
