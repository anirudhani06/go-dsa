package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 4, 8}

	result := findIndex(arr, 0, 4, []int{})
	fmt.Println(result)
}

func findIndex(arr []int, i, target int, indexes []int) []int {

	if i == len(arr)-1 {
		return indexes
	}

	if arr[i] == target {
		indexes = append(indexes, i)
	}

	return findIndex(arr, i+1, target, indexes)
}
