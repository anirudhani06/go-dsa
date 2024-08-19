package main

import "fmt"

// 1
// 23
// 456
// 78910
// 1112131415

func main() {
	n := 5
	count := 1
	for row := 1; row <= n; row++ {
		for col := 1; col <= row; col++ {
			fmt.Printf("%d", count)
			count++
		}
		fmt.Println()
	}
}
