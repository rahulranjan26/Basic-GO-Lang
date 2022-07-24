package main

import "fmt"

func main() {
	fmt.Println("Welcome to the world of GO Methods")
	rahul := User{"Rahul", 27, "rahul@gmail.com", true}
	rahul.UserIntro()

}

type User struct {
	Name     string
	Age      int
	Email    string
	IsActive bool
}

func (user User) UserIntro() {
	fmt.Printf("Hello my name is, %v. My age is %v and my email is %v", user.Name, user.Age, user.Email)
}
