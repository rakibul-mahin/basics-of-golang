package main

import(
	"fmt"
)

func main() {

	count := 1

	//This is the map
	//This is similar to dictionary and object in python and java script respectively
	var db map[string]string = map[string]string{
		"Name": "Mahin",
		"Semester": "SUMMER 2021",
		"Department": "CSE",
	}

	for cat, val := range db{
		if count == 1{
			fmt.Println(cat+" of the student: "+val+".")
		}else if count == 2{
			fmt.Println(cat+" of the student: "+val+".")
		}else{
			fmt.Println("He is in "+val+" "+cat+".")
		}

		count++
		
	}
}