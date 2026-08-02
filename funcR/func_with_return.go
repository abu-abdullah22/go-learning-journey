package main

import "fmt"

// Function that takes two integers as input and returns their sum
func return_func(num1 int, num2 int) int {
	sum2 := num1 + num2

	return sum2
}

// Function that takes two integers as input and returns their sum and product
func getNumbers(num1 int, num2 int) (int, int) {
	sum := num1 + num2

	prod := num1 * num2

	return sum, prod
}

func main() {
	// Calling the getNumbers function and printing the results
	sum, prod := getNumbers(8, 9)
	fmt.Println("Sum:", sum)
	fmt.Println("Product:", prod)

	// Calling the return_func function and printing the result
	fmt.Println(return_func(10, 20))

}
