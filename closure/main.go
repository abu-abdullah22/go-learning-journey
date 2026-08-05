package main

import "fmt"

const a = 10 // code segment

var p = 100 // run time - data segment

// code segment
func outer() func() {
	money := 100
	age := 30

	fmt.Println("Age =", age)

	// code segment - binded to outer
	show := func() {
		money = money + a + p
		fmt.Println(money)
	}

	return show // escape analysis - money is stored to heap now
}

// code segment
func call() {
	incr1 := outer()

	incr1()
	incr1()

	incr2 := outer()

	incr2()
	incr2()

}

// code segment
func main() {
	call()
}

// code segment
func init() {
	fmt.Println("hello")
}

// code segment - read only - created in the compile time

// data segment - run time creation
// executin or run time - functions (local scopes) are executed in stack frames
