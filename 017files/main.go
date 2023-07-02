package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("welcome to handeling files")
	content := "Hello from Go! (this willl go in text file)"

	file, err := os.Create("example.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}
	fmt.Println("length is ", length)
	defer file.Close()
	readFile("example.txt")

}

func readFile(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Printf("data in file is\n%v", string(data))
}
