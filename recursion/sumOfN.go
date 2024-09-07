package main

import "fmt"

func main() {
	n := 5
	result := sumOfN(n)
	fmt.Println(result)
}

func sumOfN(n int) int {
	if n <= 1 {
		return n
	}

	return n + sumOfN(n-1)
}
