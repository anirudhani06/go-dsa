package main

import "fmt"

var sum int

func main() {
	n := 1234
	reverse(n)
	fmt.Println(sum)
}

func reverse(n int) {
	if n == 0 {
		return

	}

	rem := n % 10

	sum = sum*10 + rem
	reverse(n / 10)
}
