package main

import "fmt"

// Perform division
func main() {
	var num1 int = 17
	var num2 int = 5
	var divisionResult, remain = division(num1, num2)
	fmt.Println(divisionResult)
	fmt.Printf("The amount remainfor %v %v", divisionResult, remain)
}

func division(num1 int, num2 int) (int, int) {
	var divisionResult int = num1 / num2
	var remain int = num1 % num2
	return divisionResult, remain
}
