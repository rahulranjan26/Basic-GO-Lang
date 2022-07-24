package main

//
import (
	"fmt"
	"io"
	//"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Welcome to the Web request in GO tutorial")

	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf("%T", response)

	defer response.Body.Close() //User's responsibility to close the response exclusively. In GO the response is not close on its own

	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err)

	}
	content := string(databytes)
	fmt.Println(content)
}


