package main

import "fmt"

// global scope
var (
	a = 10
	b = 20
)

// global scope
func printNum(num int) {
	// local scope
	if num > 0 {
		// block scope
		j := 100
		fmt.Println(num + j)
	} else {
		fmt.Println("Number is not positive")
	}

}

// global scope
func add(x int, y int) {
	// local scope
	res := x + y
	printNum(res)
}

// global scope
func main() {
	// local scope
	add(a, b) // a, b are global variables
}
