package main

// Use var to DECLARE three variables. The variables should have package level scope.
// Do not assign VALUES to the variables. Use the following IDENTIFIERS for the
// variables and make sure the variables are of the following TYPE (meaning they can
// store VALUES of that TYPE.

// 1. identifier "x" type int
// 2. identifier "y" type string
// 3. identifier "z" type bool

import (
	"fmt"
)

var (
	x int
	y string
	z bool
)

// print out the values of each identifier
// the compiler assigned values to the variables. What are these called values?
// r. for above "the zero value" aka "the default values"

func main() {
	fmt.Println(x, y, z)
}