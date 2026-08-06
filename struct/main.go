package main

import "fmt"

type User struct { // custom type
	Name string // member variable or property : name and age
	Age  int
}

func main() {
	var user1 User

	// instantiating an instance
	user1 = User{ // instance
		Name: "Ishtiaq",
		Age:  30,
	}

	user2 := User{ // instance
		Name: "Rowdro",
		Age:  16,
	}

	fmt.Println("user 1 : ", user1)
	fmt.Println("user 2 : ", user2)
}
