package main

import(
	"fmt"
	"strconv"
)

func main() {
	//While Loop, but with for :)
	x := 0

	for x <= 10{
		fmt.Println(x)
		x += 1
	}

	//we can add break and continue
	age := 1
	count := 1
	for age <= 13{
		if age == 13{

			if count >= 5{
				fmt.Println("Finally you are a teenager")
				break
			}else{
				fmt.Println("Nice you are a teenager")
				break
			}
			
		}

		count++
		age += 1
	}
	
	//For loop is similar to other programming languages
	for j := 0; j < 5; j++{
		fmt.Println(j)
	}

	//Lets do one multiplication project
	multiplicationOf := 13
	for n:=0; n<=20; n+=1{
		fmt.Println(strconv.Itoa(multiplicationOf)+" X "+strconv.Itoa(n)+" = "+strconv.Itoa(multiplicationOf * n))
	}

	//This same can be done with while loop
	p := 17
	m := 0
	for m <= 20{
		fmt.Println(strconv.Itoa(p)+" X "+strconv.Itoa(m)+" = "+strconv.Itoa(p * m))
		m++
	}

}