package main

import(
	"fmt"
)

func main() {
	//Switch
	num := 13

	switch num{
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	case 5:
		fmt.Println("five")
	default:
		fmt.Println("Not between 1 to 5")
	}
}