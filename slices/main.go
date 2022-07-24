package main

// We will ko about slices in this chapter

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Welcome to the slices chapter of Golang")

	// Normal arrays  ==>  Requires the information of size of array beforehand
	var fruits = [4]string{}
	fruits[0] = "apples"
	fruits[1] = "apples"
	fruits[2] = "apples"
	fruits[3] = "apples"
	fmt.Println(fruits)

	// Can't do much with such setup of array. This is static but we need dynamic

	var basket = []string{} //This is slices
	basket = append(basket, "Hello", "Mellow")
	fmt.Println(basket)

	// We can also sort the slices
	var numbers = []int{}
	numbers = append(numbers, 2, 5, 4, 9, 6, 3, 4, 2, 0, 3, 4, 5, 8)
	fmt.Println(numbers)
	// Sorting
	sort.Ints(numbers)
	fmt.Println(numbers)

	// We can also slice the slices like we do in Python
	fmt.Println(numbers[:5])
	// var myNumber uint32 = 12345678909876543212
	// fmt.Println(myNumber)

}
