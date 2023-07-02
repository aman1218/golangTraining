package main

import "fmt"

func main() {
	fmt.Println("structs example")

	john := User{1, "John", "john@example.com", 25, true}
	fmt.Println(john)
	fmt.Printf("john details are: %+v\n", john)
	fmt.Printf("john's name is: %v and john's email is: %v\n", john.Name, john.Email)
	fmt.Println("john's ID is:", john.getID())
	john.NewMail()
	fmt.Printf("john's name is: %v and john's email is: %v\n", john.Name, john.Email)
	john.NewMailPointer()
	fmt.Printf("john's name is: %v and john's email is: %v\n", john.Name, john.Email)
}

type User struct {
	ID     int
	Name   string
	Email  string
	Age    int
	Status bool
}

func (u User) getID() int {
	return u.ID
}

func (u User) NewMail() {
	u.Email = "new@mail.com"
	fmt.Println("\nnew email is:", u.Email)
}

func (u *User) NewMailPointer() {
	u.Email = "new@mail.com"
	fmt.Println("\nnew email is:", u.Email)
}