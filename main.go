package main

import "fmt"

var x = 56 // global scope

const p string = "who are you?" // Declare a constant

const q = 34 // Declare another constant

// Declare multiple constant
const (
	Pi       = 3.14
	Language = "GO"
)

func main() {

	// Declare and initialize a variable
	a := 20 // local scope
	b := "golang"
	c := 3.59
	d := true
	// i := "stored in i"  // comment out unused variable

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

	// i is not being used - invalid code
	// uncomment i and run 'go run main.go' to see what happens

	fmt.Println("p - ", p)
	fmt.Println("q - ", q)
	fmt.Println(Pi)
	fmt.Println(Language)

}
