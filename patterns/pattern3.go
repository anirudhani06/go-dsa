package main

import "fmt"

// *
// **
// ***
// ****
// *****

func main() {
	n := 5

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}
