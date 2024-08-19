package main

import "fmt"

// 55555
// 4444
// 333
// 22
// 1

func main() {
	n := 5

	for row := 1; row <= n; row++ {
		for col := 1; col <= n-(row-1); col++ {
			fmt.Printf("%d", n-(row-1))
		}
		fmt.Println()
	}
}
