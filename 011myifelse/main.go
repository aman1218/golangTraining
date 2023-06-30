package main

import "fmt"

func main() {

	var point = 7
	var result string

	if point > 7 {
		result = "good score"
	} else if point < 7 {
		result = "not bad"
	} else {
		result = "so so"
	}

	fmt.Println(result)

	if point % 2 == 0 {
		fmt.Println("num is even")
	} else {
		fmt.Println("num is odd")
	}

	if num := 8; num > 7 {
		fmt.Println("num is greater than 7")
	} else {
		fmt.Println("num is less than 7")
	}

}
