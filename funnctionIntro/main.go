package main

import "fmt"

func main() {
	fmt.Println("Welcome to the world of Functions in Golang")

	sum, message := getSum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(message, sum)


}

func getSum(numbers ...int) (int, string) {
	total := 0
	for _, val := range numbers {
		total += val
	}
	return total, "The sum is:"

}
