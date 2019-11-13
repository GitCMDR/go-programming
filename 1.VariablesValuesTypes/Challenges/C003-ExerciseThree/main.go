package main

import "fmt"

// At the package level scope, assign the following values to the three variables
// 1. for x assign 42
// 2. for y assign "James Bond"
// 3. for z assign true

var (
	x int = 42
	y string = "James Bond"
	z bool = true
)

// In func main:
// 1. Use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the
// returned value of TYPE string using the short declaration operator VARIABLE with
// the IDENTIFIER "s"

func main() {
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
}