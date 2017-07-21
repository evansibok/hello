package main

import "fmt"

var x = 56 // global scope

func main() {

	// Declare and initialize a variable

	a := 20 // local scope
	b := "golang"
	c := 3.59
	d := true

	// Print variable and typeof

	fmt.Printf("%v %T \n", a, a)
	fmt.Printf("%v %T \n", b, b)
	fmt.Printf("%v %T \n", c, c)
	fmt.Printf("%v %T \n", d, d)

	// Declare variable with zero value

	var e int
	var f string
	var g float64
	var h bool

	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", g)
	fmt.Printf("%v \n", h)

	fmt.Println()

	fmt.Println(x)

}
