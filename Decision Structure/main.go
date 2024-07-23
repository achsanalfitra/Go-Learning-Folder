package main

import "fmt"

func main() {
	var trueOrNot bool
	trueOrNot = true
	aNumber := 1

	if aNumber != 1 && trueOrNot {
		fmt.Println("All good")
	} else if trueOrNot {
		fmt.Println("Hey! only trueOrNot is true")
	} else {
		fmt.Println("Whatever about this result")
	}

	switch aNumber {
	// Switch case is like if else, but the variable declared after switch is matched with several options following the case statement
	case 1:
		// If aNumber equals 1
		fmt.Println("Now, aNumber is set to", aNumber)
	case 2:
		// If aNumber equals 2
		fmt.Println("Now, aNumber is set to", aNumber)
	case 3:
		// If aNumber equals 3
		fmt.Println("Now, aNumber is set to", aNumber)
	default:
		// if aNumber is none of the above
		fmt.Println("Now, aNumber is not here!")
	}
}
