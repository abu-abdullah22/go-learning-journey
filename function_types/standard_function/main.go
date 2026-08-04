package main

import "fmt"

// add is a standard or named function
func add(x, y int) {
	fmt.Println("sum is :", x+y)
}

// main function is the entry point of the program
func main() {
	add(4, 5)
}
