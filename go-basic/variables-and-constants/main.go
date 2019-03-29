package main

import (
	"fmt"
	"reflect"
)

func main() {
	//Explicitly declaring a "string" variable
	var explicit string = "Hello, I'm an explicitly declared variable"

	//Implicitly declaring a "string". Type inferred
	inferred := "Hello, I'm an inferred variable "

	fmt.Println(explicit)
	fmt.Println(inferred)

	fmt.Println("Variable 'explicit' is of type:",
		reflect.TypeOf(explicit))
	fmt.Println("Variable 'inferred' is of type:",
		reflect.TypeOf(inferred))
}
