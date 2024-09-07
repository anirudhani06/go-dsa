package main

import "fmt"

func main() {
	result := fibb(7)
	fmt.Println(result)
}

func fibb(number int) int {
	if number < 2 {
		return number
	}
	return fibb(number-1) + fibb(number-2)
}
