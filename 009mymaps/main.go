package main

import (
	
	"fmt"
)

func main(){

	myMaps := make(map[string]string)
	
	myMaps["JS"] = "Javascript"
	myMaps["CPP"] = "C++"
	myMaps["PY"] = "Python"
	
	fmt.Println("elements of maps are:", myMaps)
	fmt.Println("JS is short for:", myMaps["JS"])

	delete(myMaps, "CPP")
	fmt.Println("updated elements of maps are:", myMaps)
}
