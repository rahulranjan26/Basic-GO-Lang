package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the world of Goland IO"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name ")
	data, _ := reader.ReadString('\n')
	fmt.Println(data)
	fmt.Printf("The type of data is %T", data)

}
