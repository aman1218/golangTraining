package main

import "fmt"

func main() {
	var username string = "aman"
	fmt.Println(username)
	fmt.Printf("variable type %T \n\n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable type %T \n\n", isLoggedIn)

	var smallVal int = 296
	fmt.Println(smallVal)
	fmt.Printf("variable type %T \n\n", smallVal)

	var smallFloat float64 = 6969.6969
	fmt.Println(smallFloat)
	fmt.Printf("variable type %T \n\n", smallFloat)

	noVarStyle := "hello"
	fmt.Println(noVarStyle)
}
