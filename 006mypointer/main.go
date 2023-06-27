package main

import "fmt"

func main() {

	myNumber := 25

	var ptr = &myNumber

	fmt.Println("value on pointer is", *ptr)
	fmt.Println("address of pinter is", ptr)

	*ptr *= 2
	fmt.Println("updated value is", *ptr)
}