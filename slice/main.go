package main

import "fmt"

func main() {
	arr := [6]string{"this", "is", "a", "golang", "interview", "question"}
	fmt.Println(arr)

	s := arr[1:] // slice
	fmt.Println(s)

	// slice has 3 parts : 1. pointer, 2. length and 3. capacity

	s1 := s[0:2]
	fmt.Println(s1)

	fmt.Println(len(s1))
	fmt.Println(cap(s1))

	s2 := []int{3, 4, 5} // slice literal
	fmt.Println(s2, len(s2), cap(s2))

	// slice declaration with make function
	s3 := make([]int, 3)

	fmt.Println(s3, len(s3), cap(s3))

	// make function with length 3 and cap 5
	s4 := make([]int, 3, 5)
	fmt.Println(s4)

	var s5 []int // empty slice or nil slice

	fmt.Println(s5) // pointer - nil, len - 0, cap -0

}
