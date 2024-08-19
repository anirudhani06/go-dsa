package main

import "fmt"

//     *
//    ***
//   *****
//  *******
// *********

func main() {

	n := 5

	for row := 1; row <= n; row++ {
		// for col := 1; col <= n-row; col++ {
		// 	fmt.Printf(" ")
		// }

		// for col := 1; col <= (row*2)-1; col++ {
		// 	fmt.Printf("*")
		// }

		for col := 1; col <= (n+row)-1; col++ {
			if col > n-row {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}

}
