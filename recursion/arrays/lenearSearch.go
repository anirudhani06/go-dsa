package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}

	result := find(arr, 4, 0)
	fmt.Println(result)
}

func find(arr []int, target, i int) bool {
	if i == len(arr) {
		return false
	}

	return (arr[i] == target) || find(arr, target, i+1)
}
