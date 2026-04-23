package main

import "os"

func main(a,b int) {
	
     operator := os.Args[2]
	if len(os.Args) != 4 {
		return
	}
       
	if operator = "/" && b == 0 {
		println ("No division by 0")
	}
	if op == "%" && b == 0 {
	println("No modulo by 0")
	return
}
var result int

switch op {
case "+":
	result = a + b
case "-":
	result = a - b
case "*":
	result = a * b
case "/":
	result = a / b
case "%":
	result = a % b
default:
	return
}

}