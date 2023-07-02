package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("World")
	defer fmt.Println("Golang")
	println("Kloudone")
	println("Hello")
	myDefer()
}

func myDefer(){
	for i := 0; i < 5; i++ {
		defer print(i,"\n")
	}
}
