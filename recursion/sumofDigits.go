package main

import "fmt"

func main() {
	n := 1345
	result := sumofDigit(n)
	fmt.Println(result)
}

func sumofDigit(n int) int {

	if n == 0 {
		return n
	}
	return (n % 10) + sumofDigit(n/10)
}
