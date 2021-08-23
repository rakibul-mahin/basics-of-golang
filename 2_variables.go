package main

import(
	"fmt"
	"strconv"
)

//Package Level variable declaration
var shadowVar int = 26


//Important Facts:
//   1. Package Level and lowercase files in this package can access the variables
//   2. Package Level and uppercase Globally visible
//   3. Block level only available to that particular block


func main() {
	fmt.Println("----Variables----")

	//Make sure to clear all the variables declared that are not being used
	//This can throw and compile time error

	var num1 int //--- USE this when you don't need to assign any value initially
	num1 = 10
	var num2 int = 20 //--- This and the other one can also be used outside the function but num3 := 1 can not be used
	num3 := 30

	//for constant use const, const varName int = 10

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)

	fmt.Println("----Types----")

	// int, float, string, bool
	// int8 => -128 to 128, int16 => -32768 to 32768, int32 => -2147483648 to 2147483647, int64 => -9223372036854775808 to 9223372036854775807
	// uint8 => 0 to 255, uint16 => 0 to 65535, uint32 => 0 to 4294967295

	var name string = "Mohammad Rakibul Hasan Mahin"
	var ans bool = true
	var height float32 = 5.5 
	fmt.Println(ans)
	fmt.Println(height)
	fmt.Println(name)

	fmt.Println("----Shadowing----")

	//Constant values can be changed if same const is created in block level

	//Please look at the var block at package level
	fmt.Println(shadowVar)
	//Reassigning
	var shadowVar string = "Re Assigned"
	fmt.Println(shadowVar)

	//So we can re assign or create new type of package level variable
	//But when it is not the case of package level, we can't do it

	//var i int = 1
	//fmt.Println(i)
	//var i string = "try" // This will throw an erroe like ".\variables.go:47:6: i redeclared in this block previous declaration at .\variables.go:45:6"
	//fmt.Println(i)

	fmt.Println("----Conversion----")

	var i int = 32
	fmt.Printf("%v, %T\n", i, i) //This is Printf, here f stands for formating. %v == value and %T == what type

	var j float32 = float32(i)
	fmt.Printf("%v, %T\n", j, j) //This might not show decimal place, but it is really converted to a float32

	//Another Example
	var k float32 = 20.5
	fmt.Printf("%v, %T\n", k, k)

	var l int = int(k)
	fmt.Printf("%v, %T\n", l, l)

	//In this example of k and l we are converting an float to int so some data might be lost
	//Here we converted 20.5 to int and got 20, so 0.5 is lost

	//Now lets see something about string conversion
	var number int = 1
	fmt.Printf("%v, %T\n", number, number)

	var convNumberToString string = string(number)
	fmt.Printf("%v, %T\n", convNumberToString, convNumberToString)

	//We can see that it shows "" after conversion of 1, the default conversion is the uni character
	//Which means it will convert the number to the unicode symbol assigned
	//To Convert it properly, we have to import a package named "strconv"

	var properlyConvered string = strconv.Itoa(number)
	fmt.Printf("%v, %T\n", properlyConvered, properlyConvered)

	//With this we can see the correct output which is, number 1 converted to 1 with string data type

	//Now we will see the opoiste, conerting string with numbers ("1234") to int

	var stringWithNum string = "-1234"
	fmt.Printf("%v, %T\n", stringWithNum, stringWithNum)

	convToNum, err := strconv.Atoi(stringWithNum)
	//For now don't bother with the if condition
	if err == nil {
		fmt.Printf("%v, %T\n", convToNum, convToNum)
	}

	//Some other conversion techniques of strconv
	b, err := strconv.ParseBool("true")
	f, err := strconv.ParseFloat("3.1415", 64)
	z, err := strconv.ParseInt("-42", 10, 64)
	u, err := strconv.ParseUint("42", 10, 64)

	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", f, f)
	fmt.Printf("%v, %T\n", z, z)
	fmt.Printf("%v, %T\n", u, u)

	//-----------SUMMARY-----------
	// 1. Types of Variable Declaration
	// 2. Shadow Varibale
	// 3. conversion
	// 4. Visibility

}