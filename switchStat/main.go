package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to Golang for Switch")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	switch diceNumber {
	case 1:
		fmt.Printf("You got %v\n", diceNumber)

	case 2:
		fmt.Printf("You got %v\n", diceNumber)

	case 3:
		fmt.Printf("You got %v\n", diceNumber)

	case 4:
		fmt.Printf("You got %v\n", diceNumber)
		fallthrough

	case 5:
		fmt.Printf("You got %v\n", diceNumber)

	case 6:
		fmt.Printf("You got %v. Roll the dice again\n", diceNumber)
	default:
		fmt.Println("Dice game is Bad")
	}
}
