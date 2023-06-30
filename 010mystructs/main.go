package main

import "fmt"

func main() {
	fmt.Println("structs example")

	john := User{1, "John", "john@example.com", 25, true}
	fmt.Println(john)
	fmt.Printf("john details are: %+v\n", john)
	fmt.Printf("john's name is: %v and john's email is: %v\n", john.Name, john.Email)
}

type User struct {
	ID     int
	Name   string
	Email  string
	Age    int
	Status bool
}
