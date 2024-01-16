package main

import "fmt"

func main() {

	//   Single case
	day := 4

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuseday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	}

	// Multi Case
	mday := 6

	switch mday {
	case 1, 3, 5:
		fmt.Println(mday, "is Odd Day")
	case 2, 4:
		fmt.Println(mday, "is Even Day")
	case 6, 7:
		fmt.Println(mday, "is Weekend")
	default:
		fmt.Println("Invalid day no. Enter from 1 to 7")
	}

}
