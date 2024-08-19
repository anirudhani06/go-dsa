package main

import "fmt"

// 1
// 12
// 123
// 1234
// 12345

func main() {

	n := 5

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d", j)
		}
		fmt.Println()
	}

}
