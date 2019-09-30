package main

import (
	"fmt"
)

func main() {
	new_int, number := 0, 49918060
	for number > 0 {
		remainder := number % 10
		new_int *= 10
		new_int += remainder
		number /= 10
	}
	fmt.Println("Reversed Number is :", new_int)
}

/*
Reversed Number is : 6081994
*/

