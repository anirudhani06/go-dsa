package main

import "fmt"

func main() {
	digit := 0

	result := countZeros(digit, 0)
	fmt.Println(result)
}

func countZeros(digit, count int) int {
	rem := digit % 10

	if rem == 0 {
		count++
	}

	if rem == digit {
		return count
	}

	return countZeros(digit/10, count)

}
