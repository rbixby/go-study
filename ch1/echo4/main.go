/*
echo4 does the exercises from this section of the book

Exercise 1.1: Also print Args[0]
Exercise 1.2: Print the index and value of each of the arguments,
one per line
*/
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Ex 1.1: Include the command
	fmt.Println(strings.Join(os.Args[0:], " "))
	fmt.Println()

	// Ex. 1.2: Index and value, one per line
	for idx, arg := range os.Args[1:] {
		fmt.Print(idx)
		fmt.Print(":")
		fmt.Println(arg)
	}
}
