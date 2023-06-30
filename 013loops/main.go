package main

import "fmt"

func main() {
	fmt.Println("loops example")
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println("Days of the week:", days)

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}
	fmt.Println("for range type of loop example")
	for i := range days {
		fmt.Println(days[i])
	}
	fmt.Println("for each type of loop example")
	for i, day := range days {
		fmt.Printf("on index %v value is %v\n", i, day)
	}
}
