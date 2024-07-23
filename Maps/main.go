package main

import "fmt"

type Cat struct {
	Species string
	Coulour string
	Weight  float64
	Age     int
}

// For the sake of proving that maps can store anything, we create this Cat struct

func main() {
	var placeHolderOne int = 1
	var placeHolderTwo int = 2
	// Placeholder variables storing int for creating a map

	mapFromVariable := make(map[string]int)
	// This is the standard way to create a map.
	// make() is a built in function in Go to create map and other types of data structure.
	// The syntax map[string] means to create a map that takes string value as its key.
	// The syntax int is the type value of a key

	mapFromVariable["One"] = placeHolderOne
	mapFromVariable["Two"] = placeHolderTwo
	// This is the initialization of the map by giving the value of placeHolderOne and placeHolderTwo for the key "One" and "Two"

	myCat := Cat{
		Species: "Domestic",
		Age:     2,
	}
	// To create a map consisting of a struct, the struct cat is initialized with species and age

	mapFromStruct := make(map[string]Cat)
	mapFromStruct["MyCat"] = myCat
	// Similar to previous, this one stores myCat object, a type of Cat, in the key "MyCat"

	fmt.Println("I initially had", mapFromVariable["Two"], "cats. Now I only have", mapFromVariable["One"], "cat")
	fmt.Println("It is a", mapFromStruct["MyCat"].Species, "cat, aged", mapFromStruct["MyCat"].Age)
	// These print commands use all the map object created previously
	// For printing map from struct, we can specify what to print by calling the attribute of the key value
}
