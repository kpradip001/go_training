package main

import (
	"fmt"
)

func main() {
	{
		rows, stars, spaces := 4, 1, 3

		for i := 1; i < rows*2; i++ {
			for j := 1; j <= spaces; j++ {
				fmt.Print(" ")
			}
			for j := 1; j < stars*2; j++ {
				fmt.Print("*")
			}
			fmt.Println("")
			if i < rows {
				spaces--
				stars++
			} else {
				spaces++
				stars--
			}
		}
	}
}

/*
   *
  ***
 *****
*******
 *****
  ***
   *
*/

