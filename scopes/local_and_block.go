package main

import (
	"fmt"
	"example.com/mathlib"
)

/*
1. block -> {} --- anything with curly braces are blocks
and block creates local scope
2. go mod init anything.com
*/

func main(){
	x := 18 

	if x >= 18 {
		p := 10 
		fmt.Println("I am a matured boy!")
		fmt.Println("I have", p, "books.")
	}

	add(4, 7)
	mathlib.AddFromMathlib(4,7)
}