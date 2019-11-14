package main

// Continuation on Exercise Four

// At the package level scope, using the var keyword, create a VARIABLE with the IDENTIFIER
// "y". The variable should be of the underlying TYPE of your CUSTOM type "x"

import (
	"fmt"
)

type taco int
var x taco

var y int // underlying type of taco is int

// now use CONVERSION to convert the TYPE of the VALUE stored in "x" to the UNDERLYING TYPE
// then use the short declaration operator to ASSIGN that value to "y"

func main() {
	fmt.Printf("%#v\t%T\n", x, x)
	x = 42
	fmt.Println(x)

	y = int(x)
	fmt.Printf("%#v\t%T\n", y, y)
}