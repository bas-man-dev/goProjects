/* Trying to emulate a project I did for CS50 Mario -
The idea is to get feedback from the user to see how high the pyramids are
I did this with recursion on a function at one stage, I also had a while
loop protecting against user input error. Might come back to it later but currently just
trying to get the basics of go under my belt */
// even though shows an error on VSCode, compiles fine on 'go build'.

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
