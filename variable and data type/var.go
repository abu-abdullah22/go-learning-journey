package main

import "fmt"

func main(){
	printVar()
}

func printVar() {
	
	// a:=10 
	var a = 10

	// a = 10  - if we want to reassign but it has to be the same type

	var b int = 20 

	const p = 100 // we can't change value of const

	fmt.Println(a,b,p)
}

/*
data types are 
string, 
bool, 
int, int8, int16, int32, int64, 
uint, uint8, uint16, uint32, uint64, 
float32, float64
*/

