package main

import "fmt"

type Tree struct {
	// This is the format for struct declaration.
	// It is a type which consists of multiple variables of multiple types.

	NumberOfBranches int
	NumberOfRoots    int
	NumberOfLeaves   int
	Age              int
	// There are four variables in this struct, each declared using variable name then followed by variable type.
}

func main() {
	myTree := Tree{
		NumberOfBranches: 100,
		NumberOfRoots:    100,
	}
	// This is the syntax to create a variable using a struct.
	// This format uses the semicolon-equal sign and followed by the struct name with curly braces.
	// Each variable inside the struct is assigned with colon, not equal sign. And, it ended with comma for each line
	// Notice that I can assign only two out of four variables

	fmt.Println(myTree.NumberOfBranches, myTree.NumberOfRoots)
	// Printing the variable inside a struct by calling the variable name after calling the variable name, separated by period without spaces.
}
