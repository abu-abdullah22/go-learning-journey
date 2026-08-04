package main

import "fmt"

var a = 10

// add function is used to change the value of a
// functions can change the value of global variables
func add() {
	a = 26
	// printing the value of a which is changed in add function
	fmt.Println("Value of a in add function:", a)
}

func main() {
	// main function is the entry point of the program
	fmt.Println("Hello, Init function, this is main")
	// printing the value of a which is changed in init function
	add()

	fmt.Println("Value of a in main function:", a)
}

func init() {
	// init function is called before main function
	fmt.Println("Init function is called before main function and value of a is: ", a)
	// changing the value of a
	a = 20
}
