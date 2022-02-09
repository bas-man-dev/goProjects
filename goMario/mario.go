package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("hello World")

	fmt.Printf("How high do you want the pyramids?: ")
	var height int
	fmt.Scan(&height)

	for i := 1; i < (height + 1); i++ {
		gaps := strings.Repeat(" ", (height - i))
		blocks := strings.Repeat("#", i)
		fmt.Printf(gaps)
		fmt.Printf(blocks)
		fmt.Printf("  ")
		fmt.Println(blocks)
	}

}
