/* Trying to do the coin return
   things making it harder for error checking was the lack of a do while loop
   I solved this with a for loop and a break on correct format.*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var num1 float64
	for {
		fmt.Println("enter a number: ")
		fmt.Scan(&num1)
		if num1 > 0 {
			break
		}
	}

	res_1 := int(num1 * 100)
	res_2 := math.Round(res_1 * (33 / 5))
	//coins := 0

	fmt.Printf("Result is: %.2f \n", num1)

}
