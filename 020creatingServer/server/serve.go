package server

import (
	"fmt"
	"log"
	"net/http"
)

func Serving() {
	/*handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	}*/

	//http.HandleFunc("/", handler)

	fmt.Println("Server listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}