package main

import "fmt"

func main() {
	// arr := [6]string{"this", "is", "a", "golang", "interview", "question"}
	// fmt.Println(arr)

	// s := arr[1:] // slice
	// fmt.Println(s)

	// // slice has 3 parts : 1. pointer, 2. length and 3. capacity

	// s1 := s[0:2]
	// fmt.Println(s1)

	// fmt.Println(len(s1))
	// fmt.Println(cap(s1))

	// s2 := []int{3, 4, 5} // slice literal
	// fmt.Println(s2, len(s2), cap(s2))

	// // slice declaration with make function
	// s3 := make([]int, 3)

	// fmt.Println(s3, len(s3), cap(s3))

	// // make function with length 3 and cap 5
	// s4 := make([]int, 3, 5)
	// fmt.Println(s4)

	// var s5 []int // empty slice or nil slice

	// fmt.Println(s5) // pointer - nil, len - 0, cap -0

	// /*
	// ***** slice operations ******
	//  */
	// s5 = append(s5, 3) // len = 1, cap = 1

	// fmt.Println(s5)

	// var x []int

	// x = append(x, 1)
	// x = append(x, 2)
	// x = append(x, 3)

	// y := x

	// x = append(x, 4)
	// y = append(y, 5)

	// x[0] = 10

	// fmt.Println(x)
	// fmt.Println(y)

	// slice underlying array rule :
	// till 1024 - it doubles or 100%
	// after 1024 - it will be 25%^

	//// Another exmaple //////

	// z := []int{1, 2, 3, 4, 5}
	// z = append(z, 6)
	// z = append(z, 7)

	// fmt.Println(cap(z), len(z))

	// d := z[4:]
	// fmt.Println(cap(d), len(d))

	// e := changeSlice(d)

	// fmt.Println(z)
	// fmt.Println(e)

	/// variadic function
	print(5, 6, 76, 34, 3)
}

// func changeSlice(a []int) []int {
// 	a[0] = 10
// 	a = append(a, 11)
// 	return a
// }

func print(numbers ...int) {
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
}
