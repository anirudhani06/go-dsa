package main

import "fmt"

// *
// **
// ***
// ****
// *****
// ****
// ***
// **
// *

func main() {
	n := 5

	for row := 1; row <= (n*2)-1; row++ {

		var rowChange int

		if rowChange = row; row >= n {
			rowChange = (n * 2) - row
		}

		for col := 1; col <= rowChange; col++ {
			fmt.Printf("*")
		}

		fmt.Println()
	}
}
