package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://kloudone.com"

func main() {
	fmt.Println("welcome to handeling web requests")
	response, err := http.Get(url)
	checkErr(err)
	fmt.Printf("response is of type %T", response)
	//fmt.Println("response is", response)
	defer response.Body.Close()

	databytes, err := io.ReadAll(response.Body)
	checkErr(err)
	content := string(databytes)
	fmt.Println(content)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
