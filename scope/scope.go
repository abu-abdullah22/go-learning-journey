package main

import "fmt"

var a = 20
var b = 30 


func add(x int, y int){
	z := x + y 
	// z := x + a  a could be accessed here as it is global 
	// z := x + q  q couldn't be accessed as it is in block scope 
	fmt.Println(z)
}


func main(){
	var p = 30
	var q = 40 

	add(p,q)

	add(a,b )

	add(a,p)

	// add(b, z) --- z can't be accessed here 
}