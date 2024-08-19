package main

import "fmt"

// 1111
// 2222
// 3333
// 4444

func main() {
	n := 4

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Printf("%d", i)
		}
		fmt.Printf("\n")
	}

}
