package main

import(
	"fmt"
)

func main() {
	age := 13

	if age < 1{
		fmt.Println("You are a kid")
	}else if age <= 12{
		fmt.Println("wait one more year")
	}else{
		fmt.Println("I think you should sleep")
	}

	//Like other languages, we can use and, or or not
	marks := 41

	if marks < 50 && marks >= 41{
		fmt.Println("Both Condition is true")
	}else{
		fmt.Println("One or both condition is false")
	}
}