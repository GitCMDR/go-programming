package main

import "fmt"

func main() {
	fmt.Println("This is a quick example to play with basic control flow in Go")
	callingThisRandomFunction()
	fmt.Println("I'm not the function above")
}

func callingThisRandomFunction() {
	fmt.Println("You just called this random function!")
}
	// Control Flow:
	// (1) Sequence (Normally it's top to bottom)
	// (2) Loop; Iterative
	// (3) Conditional