// This program teaches about the use of pointer.
// Pointer refers to a variable which can be use to modify its value

package main

import "fmt"

func main() {
	var aString string
	aString = "Just Word"
	// The variable aString is initially defined as "Just Word"

	fmt.Println("The variable aString contains", aString)
	// Now it is being printed

	changeUsingAPointer(&aString)
	// This calls the function changeUsingAPointer() with the argument &aString
	// The ampersand symbol before the variable aString is used for pointing to an asterisks variable

	fmt.Println("After the function is modified, aString is", aString)
	// This now prints new value
}

func changeUsingAPointer(s *string) {
	// This is a function that changes the value of a string variable using a pointer
	// This function takes argument s which has a type of *string, indicating a pointer variable

	fmt.Println("The variable s is pointing towards", s)
	// This shows what s actually is, a pointer to the variable of the given argument

	var oldValue string = "Not Just Word"
	*s = oldValue
	// Declaration of the value "Not Just Word" to s
	// *s modifies the pointed value
}
