// This is my first Go program!
// The function of this program is simply to print "Hello World" in the command line interface.

package main

// This is for including the name of the file

import "fmt"

// This is automatically imported when fmt is used when using the VS code

func main() {
	// This is a main function that is a must for a Go program (if this function was empty, it is still called a function).

	// To create a variable, use the keyword "var" followed by the variable name in camel case, then ended with the type
	// Give it an equal sign to assign a value

	var aWord string = "for"
	// A variable containing the string "for"

	var aNumber int
	aNumber = 1
	// Different way to create a variable
	// This assigns the number 1 to the variable aNumber

	fmt.Println("Hello World", aWord, aNumber, "Time")
	// Printing both variable in the Println function

	var oneString string = aFunctionThatReturnsOneString()
	// One way to assign a variable to a function that returns a string

	stringOne, stringTwo := aFunctionThatReturnsTwoStrings()
	// Another way to assign a variable to a function that returns a string
	// In this case, two string outputs are given by the function, hence the necessity for two variables

	fmt.Println("The function aFunctionThatReturnsOneString() returns", oneString, "the other function returns", stringOne, "and", stringTwo)
}

func aFunctionThatReturnsOneString() string {
	// This is a function that returns one string output
	return "String One"
}

func aFunctionThatReturnsTwoStrings() (string, string) {
	// This is a fucntion that returns two string output
	return "String One", "String Two"
}
