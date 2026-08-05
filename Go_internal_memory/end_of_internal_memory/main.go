package main

import "fmt"

// code segment
const a = 10

// data segment
var p = 20

// code segment
func call() {
	// code segment binded with call
	add := func(x, y int) {
		c := x + y
		fmt.Println(c)
	}

	add(4, 5)
	add(4, 12)
}

// code segment
func main() {
	call()

	fmt.Println(a)
}

// code segment
func init() {
	fmt.Println("a is a constant!")
}

/*
2 phases : 1. compilation phase
2. execution phase



****** compile phase ******

** code segment **
a = 10
call = func() {}
add = func() {}
main = func() {}
init = func() {}





go run main.go => compile - main - ./main
go build main.go => compile - main
then we have to run the compiled file manually with - ./main
*/
