// This is the package docs
package main

import (
	"bjornkpu/go/hello/morestrings"
	"fmt"
	"github.com/google/go-cmp/cmp"
)

// This is the main class
func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
