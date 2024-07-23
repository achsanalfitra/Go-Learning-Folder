package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

// Go only has for loops that can iterate over a custom struct.
// This struct is created for that sake

func main() {
	for i := 0; i < 11; i++ {
		// This for loop iterates over an index i, the teltale sign is the semicolon
		// The semicolon separates the index initial value, end condition, and step
		// In this for loop, initial value of i is 0, it ends when i is equal to 11, then it increments by 1 for each iteration
		fmt.Println(i)
		// This simply prints i
	}

	animals := []string{"Cat", "Dog", "Horse"}
	// Let's create a slice because for loop can iterate over it

	for _, animal := range animals {
		// An index is necessary for each for loop, this one uses underscore or 'blank' because we don't care about the index
		// The second index is actually acquired from the animals variable, it uses the range keyword for it
		fmt.Println(animal)
		// Prints the second index, animal, an element of animals
	}

	var cats []Cat
	cats = append(cats, Cat{Name: "Sally", Age: 7})
	cats = append(cats, Cat{Name: "Yonny", Age: 3})
	cats = append(cats, Cat{Name: "Nada", Age: 2})
	// This initialize a slice consisting of structs

	for _, cat := range cats {
		// In similar sense to just iterating over a slice, you can do this to a slice filled with structs
		fmt.Println(cat.Name, cat.Age)
		//To print it, the attribute or field of struct can be accesssed individually
	}
}
