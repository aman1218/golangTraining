package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("enter any number")
	input, _ := reader.ReadString('\n')
	fmt.Println("your input is", input)
	incinput, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("your incremented input is", incinput+1)
	}
}
