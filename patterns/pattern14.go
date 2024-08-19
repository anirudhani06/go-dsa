package main

import "fmt"

// *********
//  *******
//   *****
//    ***
//     *

func main() {
	n := 5

	for row := n; row >= 1; row-- {
		for col := 1; col <= n-row; col++ {
			fmt.Printf(" ")
		}

		for col := 1; col <= (row*2)-1; col++ {
			fmt.Printf("*")
		}

		fmt.Println()
	}

}
