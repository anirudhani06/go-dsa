package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 4, 8}

	result := findIndexWithoutArgs(arr, 0, 4)
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

func findIndexWithoutArgs(arr []int, i, target int) []int {
	res := []int{}

	if i == len(arr)-1 {
		return res
	}
	if arr[i] == target {
		res = append(res, i)
	}

	result := findIndexWithoutArgs(arr, i+1, target)

	if len(result) != 0 {
		res = append(res, result...)
	}
	return res
}
