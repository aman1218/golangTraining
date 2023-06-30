package main

import "fmt"

func main() {
	fmt.Println("this is a function declaration example")
	greetings()
	
	result := adder(4, 5)
	fmt.Println("result is:",result)
}

func greetings() {
	fmt.Println("hello, world")
}

func adder(a int, b int) int {
	return a + b
}