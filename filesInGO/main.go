package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to the class of Files handling in GO")

	content := "Hello there. This is file handling from a text file"

	file, errors := os.Create("./myTextFile.txt")

	if errors != nil {
		fmt.Println(errors)
	}

	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Printf("Length is:%v\n", length)
	defer file.Close()
	readFile("./myTextFile.txt")

}

func readFile(fileName string) {
	data, err := os.ReadFile(fileName)
	checkNilErr(err)
	fmt.Printf("The file is :%v", string(data))

}

func checkNilErr(er error) {
	if er != nil {
		fmt.Println(er)
	}

}
