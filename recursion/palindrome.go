package main

import (
	"fmt"
	"math"
)

func main() {

	isPalindrome := isPalindrome(-1)
	fmt.Println(isPalindrome)
}

func isPalindrome(x int) bool {
	reverse := helper(x)
	return reverse == x
}

func helper(x int) int {

	rem := x % 10

	if rem == x {
		return x
	}

	return (rem * int(math.Pow(10, float64(int(math.Log10(float64(x)))+1)-1))) + helper(x/10)

}
