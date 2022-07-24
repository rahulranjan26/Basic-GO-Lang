package main

import "fmt"

func main() {
	defer fmt.Println("Line One ")
	defer fmt.Println("Line Two ")
	defer fmt.Println("Line Three ")
	fmt.Println("This is the start")
	testDefer()
}

func testDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

//Defer is nothing but a LIFO structure
