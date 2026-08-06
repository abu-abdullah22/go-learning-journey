package main

import "fmt"

func main() {
	arr := [2]int{0, 3}

	fmt.Println(arr)

	arr[1] = 6
	arr[0] = 3

	fmt.Println(arr)

	var (
		arr2 = [3]string{"I", "Love", "You"}
	)

	fmt.Println(arr2)

	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Println(a)

}
