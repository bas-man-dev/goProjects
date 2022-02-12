package main

import (
	"fmt"
)

func main() {
	fmt.Println("WELCOME TO THE QUIZ!")

	fmt.Printf("Enter your name: ")

	var name1 string
	fmt.Scan(&name1)

	var age int

	for {
		fmt.Println("How old are you? ")
		fmt.Scan(&age)
		if age >= 18 {
			break
		}
		fmt.Println("you must be 18 or over!")
	}

	fmt.Printf("hi %v you are %v years old\n", name1, age)

}
