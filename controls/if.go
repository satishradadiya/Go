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
	// var num int = 9
	if num := 9; num < 3 {
		fmt.Println(num, "is negative")
	} else if num < 11 {
		fmt.Println(num, "is positive")
	} else {
		fmt.Println(num, "is something")
	}
}
