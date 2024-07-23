package main

import (
	"fmt"
	"sort"
)

// Let's import the sort module to sort slices with ints

func main() {
	var stringSlice []string
	// This is the standard way to create a slice.
	// The syntax square brackets before the string declaration differentiates between normal variable and slices

	stringSlice = append(stringSlice, "Hey!", "Why", "do", "you", "care")
	// To add elements into a slice, use the builtin append function, which takes the slice name in the first positional argument.

	intSlice := []int{1, 2, 3, 4, 5, 6, 9, 8, 7}
	// This does the same thing as the previous syntax but the shorthand version with colon-equal sign
	// This syntax does not require the function append, instead it uses curly brackets after the variable type declaration which consists of the elements desired

	sort.Ints(intSlice)
	// This is a neat function from sort to sort int slices

	fmt.Println(stringSlice[0])
	// This only prints the zeroth index element of stringSlice, which is "Hey!"

	fmt.Println(stringSlice, "that the numbers are", intSlice[0:9])
	// This prints both stringSlice and the sorted intSlice
}
