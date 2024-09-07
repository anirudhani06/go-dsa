package main

import "fmt"

func main() {
	n := 10

	oneToN(n)
}

func nTo1(n int) int {

	if n == 0 {
		return n
	}

	fmt.Println(n)

	return nTo1(n - 1)
}

func oneToN(n int) {
	if n == 0 {
		return
	}
	oneToN(n - 1)
	fmt.Println(n)

}
