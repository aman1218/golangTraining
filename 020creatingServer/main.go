package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"creatingServer/server"

)

const url = "http://localhost:3000/get"

func main() {
	fmt.Println("server and get, post")
	server.Serving()
	get(url)
	
}

func get(myurl string) {

	response, err := http.Get(myurl)
	checkError(err)
	defer response.Body.Close()
	fmt.Println("status code:", response.StatusCode)
	fmt.Println("content length:", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	//byteCount, _ := responseString.Write(content)
	fmt.Println(responseString.String())
	fmt.Println("content:", string(content))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
