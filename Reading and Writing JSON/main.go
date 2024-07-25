package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string  `json:"name"`
	Age  float64 `json:"age"`
}

// This is a struct, created for handling json objects
// This struct has a field tag after field type noted with backquote
// The format json:"name" means it handles json with "name" tag

func main() {
	myJSON :=
		`[
	{
		"name":"adam",
		"age":12
	},
	{
		"name":"aster",
		"age":15
	}
]` // This is a dummy json, it consists of two jsons with the field name and age for each object

	var myJSONInVariable []Person
	err := json.Unmarshal([]byte(myJSON), &myJSONInVariable)
	if err != nil {
		fmt.Println("Error unmarshalling your JSON", err)
	} else {
		fmt.Println(myJSONInVariable)
	}

	// This is for creating an interface object using the dummy json
	// The function used for it is json.Unmarshal
	// The function works by converting the json object into an interface object
	// This function returns an error if it exists

	var sliceForJSON []Person
	// Now to create similar json from struct or an interface object is by declaring the slice first

	var p1 Person
	p1.Name = "Andre"
	p1.Age = 15

	var p2 Person
	p2.Name = "Lily"
	p2.Age = 16

	// These two Person structs will populate the slice initially created

	sliceForJSON = append(sliceForJSON, p1, p2)

	JSONOutput, err := json.MarshalIndent(sliceForJSON, "", "     ")
	if err != nil {
		fmt.Println("Error creating JSON", err)
	} else {
		fmt.Println(string(JSONOutput))
	}
	// Similar to the previous function, this function works the other way round
	// It converts structs into json with indent due to the function MarhsalIndent
	// It takes prefixes and indent
	// Like the previous function, this function also returns error alongside with the json object
}
