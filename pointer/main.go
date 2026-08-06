package main

import "fmt"

func print(numbers *[3]int) { // pass by reference
	fmt.Println(numbers)
}
func print2(numbers [2]int) { // pass by value
	fmt.Println(numbers)
}

type User struct {
	Name string
	Age  int
}

func printUser(usr *User) {
	fmt.Println(usr)
}
func main() {
	x := 20

	a := &x
	*a = 21

	fmt.Println(x)

	arr := [3]int{2, 45, 5}
	print(&arr) // pass by reference

	arr2 := [2]int{23, 334}
	print2(arr2)

	user1 := User{
		Name: "ishti",
		Age:  24,
	}

	printUser(&user1)

}
