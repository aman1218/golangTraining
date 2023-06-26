package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("enter any value: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("your input is", input)
	fmt.Printf("type of input is %T", input)
}
