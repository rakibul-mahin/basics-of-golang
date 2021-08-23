package main

import(
	"fmt"
)

func main() {
	//declaring an array
	var arr [5]int //We have to declare how many elements we want to store
	anotherArr := [5]int{1,2,3,4,5}
	fmt.Println(arr)
	fmt.Println(anotherArr)

	//lets add all numbers in our anotherArr
	sum := 0
	for i := 0; i < len(anotherArr); i++{
		sum += anotherArr[i]
	}

	fmt.Println(sum)
}