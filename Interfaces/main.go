package main

import "fmt"

type Tree interface {
	Name() string
	BarkColour() string
	LeavesColour() string
}

// This is an interface filled with functions that can be used to give functionality to structs

type Cedar struct {
	Age int
}

// A normal struct

type Walnut struct {
	Edible string
}

// A normal struct

func main() {
	myCedar := Cedar{
		Age: 30,
	}

	myWalnut := Walnut{
		Edible: "Yes, it is edible",
	}

	aboutTree(&myCedar)
	aboutTree(&myWalnut)
	// Since the aboutTree uses the functions in the interfaces, the struct passed in here must have those functions
	// Notice the ampersand, this is because the receivers syntax used to create interface functions inside the struct
	// If the written functions for the structs do not use receiver--the asterisk symbol--then it would not be necessary to use ampersand
	// However, the use receiver and ampersand syntax is best practice
}

func aboutTree(t Tree) {
	fmt.Println("The", t.Name(), "is", t.BarkColour(), "and its leaves are", t.LeavesColour())
	// This is a function that takes the Tree interface and uses functions within it
	// All the functions here belong to the t argument passed
}

func (n *Cedar) Name() string {
	return "Cedar"
}

// This is a receiver function to note that this function belongs to Cedar
// Since now Cedar has Name() function, the interface can call this function from the Cedar struct

func (b *Cedar) BarkColour() string {
	return "White"
}

func (l *Cedar) LeavesColour() string {
	return "Light Green"
}

func (n *Walnut) Name() string {
	return "Walnut"
}

func (b *Walnut) BarkColour() string {
	return "Brown"
}

func (c *Walnut) LeavesColour() string {
	return "Dark Green"
}
