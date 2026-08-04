package main

import "fmt"

// main function
func main() {
	// immediately invoked function expression (IIFE)
	func(x, y int) {
		fmt.Println("Hello, I am an anonymous function!")
		fmt.Println("Sum of x and y is:", x+y)
	}(6, 7) // passing arguments to the anonymous function
}
