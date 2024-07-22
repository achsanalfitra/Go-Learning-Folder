// This is my first Go program!
// The function of this program is simply to print "Hello World" in the command line interface.

package main

// This is for including the name of the file

import "fmt"

// This is automatically imported when fmt is used when using the VS code

func main() {
	// This is a main function that is a must for a Go program (if this function was empty, it is still called a function).
	fmt.Println("Hello World")
	// The "fmt" stands for 'format' where we use the function Println to print "Hello World"
}
