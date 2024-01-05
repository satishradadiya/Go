package main

import "fmt"

func main() {
	var int1 int = 11
	var sum int = int1 + 6
	fmt.Println(int1)
	fmt.Println(sum)

	var flo1 float32 = 99.99
	fmt.Println(flo1)

	var sumResult float32 = float32(int1) + flo1
	fmt.Println(sumResult)

	var mystrn string = "I am string"
	fmt.Println(mystrn)

	var mybool bool = false
	fmt.Println(mybool)

	const pi float32 = 3.14
	fmt.Println(pi)

}
