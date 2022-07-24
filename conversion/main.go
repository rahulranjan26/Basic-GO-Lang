package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Please rate our pizza in between 1 and 5."
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	data, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {

	} else {
		fmt.Printf("The revised rating for pizza is:%v", data+1)
	}
}
