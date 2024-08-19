package main

import "fmt"

// *****
// ****
// ***
// **
// *

func main() {
	n := 5

	for row := 1; row <= n; row++ {
		for col := 1; col <= n-(row-1); col++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}
