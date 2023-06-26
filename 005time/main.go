package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println("current time is:", presentTime.Format("02-01-2006 15:04:05 Monday"))

	createdDate := time.Date(2023, time.June, 28, 18, 28, 27, 0, time.UTC)
	fmt.Println("created date: ", createdDate.Format("02-01-2006 15:04:05 Monday"))
}
