package main

import "fmt"

var a int
type orc int
var b orc

func main() {
	a = 42
	b = 48

	// Convert my orc type to int type so that it can be assigned to the a variable
	a = int(b)
	fmt.Println(a, b)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", a)
}