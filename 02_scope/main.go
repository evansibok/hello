package main

import "fmt"

var x = 56 // variable declared as global scope

func main() {
	a := 20 // variable declared as local scope

	// Print local scope variable
	fmt.Printf("%v %T \n", a, a)

	// Print global scope variable
	fmt.Printf("%v %T \n", x, x)

}
