package main

import (
	"fmt"

	"github.com/achsanalfitra/Go-Learning-Folder/Package/helpers"
	// This is for importing the package from specified location, best practice is the repository where it is stored
)

func main() {
	myCat := helpers.Cat{
		Name: "Sandra",
		Age:  2,
	}
	// This declares that myCat is a type of Cat from the helpers package

	fmt.Println("I have a cat, her name is", myCat.Name, "her age is", myCat.Age)
}
