package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func PrintUserData(usr User) {
	fmt.Println(usr.Name)
	fmt.Println(usr.Age)
}

// receiver function
func (usr User) PrintData() {
	fmt.Println(usr.Name)
	fmt.Println(usr.Age)
}

func (usr User) call(a int) {
	fmt.Println(usr.Name)
	fmt.Println(a)
}

func main() {
	user1 := User{
		Name: "Habib",
		Age:  30,
	}

	PrintUserData(user1)

	user2 := User{
		Name: "something",
		Age:  23,
	}

	PrintUserData(user2)

	user1.PrintData()
	user2.PrintData()

	user1.call(45)
}
