package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("switch case example")
	rand.Seed(time.Now().UnixNano())
	var randomValue int = rand.Intn(6) + 1
	fmt.Println("random value is", randomValue)


	switch randomValue {
	case 1:
		fmt.Println("your dice value is 1 you can open")
	case 2:
		fmt.Println("your dice value is 2 you can move 2 step")
	case 3:
		fmt.Println("your dice value is 3 you can move 3 step")
		fallthrough
	case 4:
		fmt.Println("your dice value is 4 you can move 4 step")
	case 5:
		fmt.Println("your dice value is 5 you can move 5 step")
		fallthrough
	case 6:
		fmt.Println("your dice value is 6 you can move 6 step and roll dice again")
	default:
		fmt.Println("invalid")
	}
}
