package main

import "fmt"

// Function that takes two integers as input and prints their sum
func fnc(num1 int, num2 int) {

	sum := num1 + num2

	fmt.Println(sum)

}

func main() {
	// Calling the fnc function with two integers
	fnc(8, 10)
}
