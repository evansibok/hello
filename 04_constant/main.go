package main

import "fmt"

const p string = "who are you?" // Declare a constant

// Declare multiple constant
const (
	Pi       = 3.14
	Language = "GO"
)

func main() {
	const q = 34 // Declare another constant

	// Print Constant
	fmt.Println("p - ", p)
	fmt.Println("q - ", q)

	fmt.Println(Pi)
	fmt.Println(Language)
}
