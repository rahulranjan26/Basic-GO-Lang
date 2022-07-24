package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Golang version of control statements")

	fmt.Println("Enter the number")
	number := bufio.NewReader(os.Stdin)
	data, _ := number.ReadString('\n')
	num,_ := strconv.ParseFloat(strings.TrimSpace(data),64)
	fmt.Printf("%T\n", num)
	//number := strconv.FormatInt(number,64)
	if num<10{
		fmt.Println("number is less than 10")
	} else if num>10{
		fmt.Println("Number is greater than 10")
	} else{
		fmt.Println("Number is equal to 10")
	}

	if tm:=3; tm<10{
		fmt.Println("Bad Number")
	}
}
