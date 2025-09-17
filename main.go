package main 

import "fmt"

func main() {
	a := 10 
	b := 20
	fmt.Println("Hello World!")

	printVar()
	ifelse_switch()
	fnc(12, 10)
	fnc(a, b)

	sum2 := return_func(10,20)
	fmt.Println(sum2)

	sum, prod := getNumbers(a,b)
	fmt.Println(sum, prod)
}
