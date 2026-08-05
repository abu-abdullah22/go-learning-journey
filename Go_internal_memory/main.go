package main

import "fmt"

// internal memory is divided into 4 parts : code segment, data segment, heap and stack

// variables are stored in the data segment
var a = 10

// functions are stored in the code segment
func add(a, b int) {
	c := a + b
	fmt.Println("Sum:", c)
}

// add is stored in the code segment and it is executed in the stack segment. stack frames are created for each function call and destroyed when the function returns.

// stored in code segment
func main() {
	fmt.Println("Main!")
	add(3, 5)
	add(5, a)
}

// after init is executed, main is executed - it happens in the stack segment - stack frames are created for each function call and destroyed when the function returns.

// stored in code segment
func init() {
	fmt.Println("Init!")
}

// when init is being executed and it is executed very first - it happens in the stack segment - stack frames are created for each function call and destroyed when the function returns. The stack segment is used for static memory allocation and is managed by the compiler.
