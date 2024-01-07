package main

import "fmt"

func main() {
	// variable and simple arothimatic example
	var int1 int = 11
	var sum int = int1 + 6
	fmt.Println(int1)
	fmt.Println(sum)

	// variable with float values
	var flo1 float32 = 99.99
	fmt.Println(flo1)

	//init value converstion to float and perform arithimatic
	var sumResult float32 = float32(int1) + flo1
	fmt.Println(sumResult)

	// string example
	var mystrn string = "I am string"
	fmt.Println(mystrn)

	// boolean example
	var mybool bool = false
	fmt.Println(mybool)

	// const ecample (unmutable variable)
	const pi float32 = 3.14
	fmt.Println(pi)

}
