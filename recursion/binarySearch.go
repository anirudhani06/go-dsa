package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 10
	result := binarySearch(arr, target, 0, len(arr)-1)

	fmt.Println(result)
}

func binarySearch(arr []int, target, start, end int) int {
	if start > end {
		return -1
	}

	mid := start + (end-start)/2

	if arr[mid] == target {
		return mid
	}
	if target > arr[mid] {
		return binarySearch(arr, target, mid+1, end)
	}
	return binarySearch(arr, target, start, mid-1)

}
