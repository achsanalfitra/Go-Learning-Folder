package main

import (
	"fmt"

	"github.com/achsanalfitra/Go-Learning-Folder/Package/helpers"
)

func main() {
	myCat := helpers.Cat{
		Name: "Sandra",
		Age:  2,
	}

	fmt.Println("I have a cat, her name is", myCat.Name, "her age is", myCat.Age)
}
