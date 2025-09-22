package main

import "fmt"

/*
1. block -> {} --- anything with curly braces are blocks
and block creates local scope
*/

func main(){
	x := 18 

	if x >= 18 {
		p := 10 
		fmt.Println("I am a matured boy!")
		fmt.Println("I have", p, "books.")
	}
}