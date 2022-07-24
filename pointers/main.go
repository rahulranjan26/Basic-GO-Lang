package main

import "fmt"

func main() {

	fmt.Println("Welcome to the dreaded world of Pointers:")

	var myNumber int = 23
	var ptr *int = &myNumber
	// ptr = &myNumber

	// var ptr *int  ==> Pointer ptr will points to int type variable
	fmt.Println("The value of pointer is :", ptr)
	*ptr = 12 //We can change the value at address pointed by the variable ptr
	fmt.Println(myNumber)

}
