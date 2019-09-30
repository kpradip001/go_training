package main

import (
	"fmt"
)

func main() {
	a, b := 1, 1
	c := a + b
	for i := 0; i <= 10; i++ {
		fmt.Print(a," ")
		c = a+b
		a, b = b, c
	}
}

/*
1 1 2 3 5 8 13 21 34 55 89 
*/

