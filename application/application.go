package main

import "fmt"


func printWelcome() {
	//welcome message
	fmt.Println("Welcome to the application!")
}

func nameInput() string {
	//Taking user input and returning the value
	var name string
	fmt.Println("Enter your name -")
	fmt.Scanln(&name)

	return name
}

func numsInput() int {
	//taking nums input and returning the value
	var num1 int
	var num2 int
	fmt.Println("Enter first number -")
	fmt.Scanln(&num1)
	fmt.Println("Enter second number -")
	fmt.Scanln(&num2)

	sum := num1 + num2

	return sum
}

func display(name string, sum int) {
	// Displaying results
	fmt.Println("Hello, ", name)
	fmt.Println("The sum of your numbers is : ", sum)

	// Goodbye message
	fmt.Println("Thank you for using the application, ", name)
	fmt.Println("Goodbye!")
}


func application() {
	// storing all the information

	printWelcome()

	name := nameInput()

	sum := numsInput()

	display(name, sum)

}


func main() {
	//calling the whole application at once
	application()
}
