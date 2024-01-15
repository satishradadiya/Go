package main

import "fmt"

func main() {
	if 6%2 == 0 {
		fmt.Println("It's even")
	} else {
		fmt.Println("It's odd")
	}
	if 8/2 == 4 || 7%2 == 0 {
		fmt.Println("One of the fiven condition matched")
	}

}
