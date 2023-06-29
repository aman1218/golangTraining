package main

import (
	"fmt"
	"sort"
)

func main() {

	var myLists = []int{11, 12, 13, 14, 15, 16, 17}
	fmt.Println("before appending:",myLists)

	myLists = append(myLists, 18, 19)
	fmt.Println("after appending:", myLists)

	myLists = append(myLists[2:7])
	fmt.Println("slicing the slices:",myLists)

	myMake := make([]int, 5)

	myMake[0] = 555
	myMake[1] = 444
	myMake[2] = 999
	myMake[3] = 333
	myMake[4] = 111
	fmt.Println("slice elements:", myMake)

	myMake = append(myMake, 666, 875, 154, 625, 561)
	fmt.Println("slice after appending:", myMake)

	sort.Ints(myMake)
	fmt.Println("after sorting:", myMake)
	
	index := 2
	myMake = append(myMake[:index], myMake[index+1:]...)
	fmt.Println("after removing index 2:", myMake)

}
