package main

import "fmt"

var a = 10

func main() {
	age := 10

	if age > 18 {
		// This will create a new variable 'a' that shadows the outer 'a'
		a := 20
		fmt.Println("Value of A becomes :", a)
	}

	// This will print the value of the outer 'a' since the inner 'a' is out of scope
	fmt.Println("Value of A is: ", a)
}
