package main

import "fmt"

var OuterVariable string = "Variable from the outerspace!"

// This is a variable declared outside all function scopes.
// This should run everywhere inside a function.

func main() {
	replacesOutsideInstead("this")
	thisPrintsOutside("This")

}

func replacesOutsideInstead(OuterVariable string) {
	fmt.Println("This function doesn't work!, it replaces the OuterVariable with", OuterVariable)
	// This is not a good function because its argument overwrites the global variable
}

func thisPrintsOutside(iDontCare string) {
	fmt.Println("This function works because it prints", OuterVariable)
	iDontCare = "whatever"
	// This is a good function because it takes the global variable, instead of overwriting it
}
