package main

import "fmt"

func application() {
	//welcome message 
	fmt.Println("Welcome to the application!")

	//Taking ser input 
	var name string 
	fmt.Println("Enter your name -")
	fmt.Scanln(&name)

	var num1 int 
	var num2 int 
	fmt.Println("Enter first number -")
	fmt.Scanln(&num1)
	fmt.Println("Enter second number -")
	fmt.Scanln(&num2)

	sum := num1 + num2

	// Displaying results 
	fmt.Println("Hello, ", name)
	fmt.Println("The sum of your numbers is : ", sum)

	// Goodbye message 
	fmt.Println("Thank you for using the application, ", name)
	fmt.Println("Goodbye!")


}

func main() {
	application()
}