package main

import "fmt"

func main() {
	fmt.Println("Welcome to map in Golang")

	languages := make(map[string]string)

	languages["js"] = "javascript"
	languages["py"] = "Python"
	languages["go"] = "Golang"
	fmt.Println(languages)

	// laguage:=[]string{}

	// Now if you have to delete
	delete(languages, "js")
	fmt.Println(languages)
}
