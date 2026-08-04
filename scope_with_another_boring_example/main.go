package main

import "fmt"

var (
	a = 10
	b = 20
)

func PrintSum(x int) {
	fmt.Println(x)
}

func add(a, b int) {
	res := a + b
	PrintSum(res)
}

func main() {
	add(a, b)
}
