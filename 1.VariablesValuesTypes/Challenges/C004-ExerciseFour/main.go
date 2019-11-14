package main

import (
	"fmt"
)

// For this exercise:
// 1. Create your own type, have the underlying type be an int.
// 2. Create a variable of your new type with the identifier "x" using the "var"
//    keyword

type taco int
var x taco

// In func main:
// Print out the value of the variable "x"
// Print out the type of the variable "x"
// Assign 42 to the variable "x" using the "="operator
// Print out the value of the variable "x"

func main() {
	fmt.Printf("%#v\t%T\n", x, x)
	x = 42
	fmt.Println(x)
}