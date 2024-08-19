package main

import "fmt"

//     *
//    **
//   ***
//  ****
// *****

func main() {
	n := 5

	for row := 1; row <= n; row++ {
		for col := 1; col <= n; col++ {
			if col <= n-row {
				fmt.Printf(" ")
			} else {
				fmt.Printf("*")
			}
		}
		fmt.Println()
	}
}
