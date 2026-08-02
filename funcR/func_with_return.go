package main

import "fmt"

func return_func(num1 int, num2 int) (int, int) {
	sum2 := num1 + num2
	sum3 := num1 * num2

	return sum2, sum3
}

func getNumbers(num1 int, num2 int) (int, int) {
	sum := num1 + num2

	prod := num1 * num2

	return sum, prod
}

func main() {
	sum, product := return_func(10, 20)
	fmt.Printf("The sum and product of 10 and 20 is: %d and %d\n", sum, product)
}
