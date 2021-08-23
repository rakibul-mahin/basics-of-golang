package main

import(
	"fmt"
	"strconv"
)

func main() {
	//Why slices??
	//This don't need to enter how big our array is

	arr1 := [5]int{1,2,3,4,5}
	var arr2 []int = arr1[1:3]
	fmt.Println(arr1)
	fmt.Println(arr2)
	/*
		a[:] ==> Include all from begining to end
		a[1:] ==> Start at 1 and go to end
		a[1:4] ==> start at 1 and dont include 4
	*/

	//Declaring Slice
	var arr3 []int = []int{1,2,3,4,5,6,7,8,9} //This created a slice without any limit
	//putting value at end of a slice
	arr3 = append(arr3, 10)

	fmt.Println(arr3)

	var footballerArray []string = []string{"L. Messi","C. Ronaldo","N. Barella","C. Immobile","L. Insigne"}
	//Traditionally we will use a for loop and access each names as below
	for i := 0; i < len(footballerArray); i++{
		fmt.Println(footballerArray[i])
	}

	//but there is another way which is effecient sometimes, this method is known as range
	for index, elem := range footballerArray{ //This line will return two variable, one the index and another what is in that index
		fmt.Println("in index "+strconv.Itoa(index)+" we have "+elem)
	}
	//The above code is similar to for i in array: in python
	//Now we know any variable declared in go have to be used but sometimes we don't want the index to be used at that time we can use an _ instead of index

}