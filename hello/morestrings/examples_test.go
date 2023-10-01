package morestrings_test

import (
	"bjornkpu/go/hello/morestrings"
	"fmt"
)

// Prints the reverse out to the console.
func ExampleReverseRunes() {
	inputString := "!oG ,olleH"
	result := morestrings.ReverseRunes(inputString)
	fmt.Println(result)
	// OutPut:
	// Hello, Go!
}
