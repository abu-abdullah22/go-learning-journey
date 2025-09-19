package main

import "fmt"

func ifelse_switch() {
	age := 18
	gender := "male"

	if age > 18 && gender == "male" {
		fmt.Println("You're eligible to marry.")
	} else if age < 18 {
		fmt.Println("You are not eligible to marry, but you can love someone")
	} else {
		fmt.Println("You are just a teenager")
	}

	// >, >=, <, <=, ==
	// && -and
	// || -or
	// not - !

	isPretty := false

	if !isPretty {
		fmt.Println("too much good looking!")
	}

	// switch
	a := 5

	switch a {
	case 1:
		fmt.Println("a is 1")
	case 5:
		fmt.Println("a is 5")
	default:
		fmt.Println("a is what?")
	}

}

func main(){
	ifelse_switch()
}
