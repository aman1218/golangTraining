package main

import (
	"fmt"
	"net/url"
)

const myURL string = "https://www.google.com/search?q=golang"

func main() {
	fmt.Println("handelingURL's")
	fmt.Println(myURL)
	result, _ := url.Parse(myURL)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	partsofURL := &url.URL{
		Scheme:   "https",
		Host:     "www.google.com",
		Path:     "/search",
		RawQuery: "q=golang",
	}
	nextURL := partsofURL.String()
	fmt.Println("nextURL: ", nextURL)

}
