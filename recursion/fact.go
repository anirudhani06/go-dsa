package main

import "fmt"

func main() {
	n := 5
	result := fact(n)
	fmt.Println(result)
}

func fact(n int) int {

	if n <= 1 {
		return 1
	}

	return n * fact(n-1)
}
