package main

import "fmt"

// 12345
// 1234
// 123
// 12
// 1

func main() {
	n := 5

	for row := 1; row <= n; row++ {
		for col := 1; col <= n-(row-1); col++ {
			fmt.Printf("%d", col)
		}
		fmt.Println()
	}
}
