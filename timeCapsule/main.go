package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Welcome to time capsule of GoLang")
	presentTime := time.Now()
	fmt.Println(presentTime)
	// This format is inbuilt. Cant do much about it.
	newTime := presentTime.Format("02-01-2006 Monday")
	fmt.Println(newTime)

	newDate := time.Date(2022, time.July, 12, 15, 30, 25, 254, time.UTC)
	fmt.Println(newDate)

}
