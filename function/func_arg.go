package main

import "fmt"

//function parameter defining and passing parameter to other function

func main() {
	var firstString string = "First string value"
	secondFunc(firstString)
}

func secondFunc(firstString string) {
	fmt.Println(firstString)
}
