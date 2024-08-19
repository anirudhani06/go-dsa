package main

import "fmt"

// 1
// 01
// 101
// 0101
// 10101

func main() {
	n := 5
	value := 1
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			value = 0
		} else {
			value = 1
		}
		for j := 1; j <= i; j++ {
			fmt.Printf("%d", value)
			if value == 0 {
				value = 1
			} else {
				value = 0
			}
		}
		fmt.Println()
	}
}
