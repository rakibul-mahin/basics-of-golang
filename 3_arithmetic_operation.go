package main

import(
	"fmt"
)

func main() {
	fmt.Println("----Arithmetic Operations----")

	//First lets declare some variables
	a := 28.3
	b := 3.3

	//Now lets do some operations
	sum := a + float64(b)
	fmt.Println(sum)
	sub := a - float64(b)
	fmt.Println(sub)
	mul := a * float64(b)
	fmt.Println(mul)
	div := a / float64(b)
	fmt.Println(div)

	//mod := a % float64(b)
	//fmt.Println(mod)
	//Not getting why we cant find remainder of two floating numbers

	//BTW even if we have one var as int(2) and the other one as int32(4) we can not add or do any operations

	//Now here in div, if the result was something like 3.25 we would get only 3

	fmt.Println("----Bit Operator----")

	c := 10
	d := 3

	andOpp := c & d // where the value matches it is 1
	fmt.Println(andOpp)
	orOpp := c | d // when there is atleast one 1 it is 1
	fmt.Println(orOpp)
	exOr := c ^ d // When there will be exactly one 1
	fmt.Println(exOr)
	andNot := c &^ d // opposite of orOpp
	fmt.Println(andNot)

	fmt.Println("----Complex----")

	comNum := 1 + 2i
	fmt.Printf("%v, %T\n",comNum, comNum)
	//There is another way using complex()
	anoComp := complex(1,2)
	fmt.Printf("%v, %T\n", anoComp, anoComp)
	
}