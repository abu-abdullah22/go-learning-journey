package main

import "fmt"

/*
features of a higher order function:
1. It can take one or more functions as arguments.
2. It can return a function as a result.
3. It can be assigned to a variable.
*/

func add(a, b int) {
	c := a + b
	fmt.Println("Sum:", c)
}

func product(x, y int) {
	z := x * y
	fmt.Println("Product:", z)
}

func processOperation(a, b int, op func(x, y int)) {
	op(a, b)
}

func main() {
	// Passing a function as an argument to another function
	processOperation(5, 3, add)
	processOperation(5, 3, product)

	// Returning a function from another function
	call()(5, 5)

	// Assigning a function to a variable
	sum := call()
	sum(10, 20)
}

// Returning a function from another function
func call() func(x, y int) {
	return add
}
