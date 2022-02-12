package main

import (
	"fmt"
)

func main() {
	fmt.Println("WELCOME TO THE QUIZ!")

	fmt.Printf("Enter your name: ")

	var name1 string
	fmt.Scan(&name1)

	fmt.Printf("How old are you? ")

	var age int
	fmt.Scan(&age)

	fmt.Printf("hi %v you are %v years old\n", name1, age)

}
