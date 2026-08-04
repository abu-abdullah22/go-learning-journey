package main

import "fmt"

func add(x, y int) {
	c := x * y
	fmt.Println("product of x and y is:", c)
}

func main() {
	add(3, 4)
	z := 5
	// anonymous function assigned to a variable
	add := func(a, b int) {
		c := a + b + z
		fmt.Println("Sum of a, b and z is:", c)
	} // expression ends here
	add(6, 7)
}
