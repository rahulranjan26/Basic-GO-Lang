package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://www.rahul.dev:3000/learn?coursename=dsa&fees=5000"

func main() {
	fmt.Println("Welcome to the GO tuts of handling URLs")

	resources, err := url.Parse(myUrl)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resources.Scheme)
	fmt.Println(resources.Path)
	fmt.Println(resources.Host)
	fmt.Println(resources.RawQuery)
	fmt.Println(resources.Port())
	queryParams := resources.Query()
	for key, val := range queryParams {
		fmt.Println(key, "=====>", val)
	}

}
