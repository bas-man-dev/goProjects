package main

import "fmt"

func main() {
	fmt.Println("WELCOME TO THE QUIZ GAME")

	fmt.Printf("Enter your name: ")
	var name string

	fmt.Scan(&name)

	fmt.Printf("Hello, %v, welcome to the game \n", name)

}
