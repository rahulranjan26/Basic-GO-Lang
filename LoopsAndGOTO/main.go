package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Golang for Loops")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Satuday"}

	if len(days) == 7 {
		goto rahul
	}

	//Way 1
	//for i := 0; i < len(days); i++ {
	//	fmt.Println(days[i])
	//}

	//	Way 2
	//for i := range days {
	//	fmt.Println(days[i])
	//}

	//	Way 3 ===> Most Preferred way
	//for index, data := range days {
	//	fmt.Printf("Day number :%v, Day:%v\n", index, data)
	//}

rahul:
	fmt.Println("Welcome to Golang classes")
}
