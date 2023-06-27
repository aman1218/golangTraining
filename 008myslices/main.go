package main

import "fmt"

func main() {

	var myLists = []int{11, 12, 13, 14, 15, 16, 17}
	fmt.Println("before appending:",myLists)

	myLists = append(myLists, 18, 19)
	fmt.Println("after appending:", myLists)

	myLists = append(myLists[2:7])
	fmt.Println("slicing the slices:",myLists)

}
