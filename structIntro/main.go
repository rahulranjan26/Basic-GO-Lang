package main

import "fmt"

type Users struct {
	Name  string
	Age   int
	Email string
	desig string
}

func main() {
	fmt.Println("Struct in Golang")

	rahul := Users{"Rahul", 27, "go@gmail.com", "SDE1"}
	fmt.Printf("%v\n",rahul)
	fmt.Printf("%+v",rahul)


}
