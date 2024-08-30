package main

import "fmt"

func main() {

	result := search([]int{1, 4, 5, 7, 12, 15, 42, 87, 90, 100}, 90)
	fmt.Println(result)

}

func search(arr []int, target int) int {
	startIndex := 0
	endIndex := 1

	for target > arr[endIndex] {
		temp := endIndex + 1
		endIndex = endIndex + (endIndex-startIndex+1)*2
		startIndex = temp

	}

	return binarySearch(arr, startIndex, endIndex, target)
}

func binarySearch(arr []int, start, end, target int) int {
	for start <= end {
		mid := start + (end-start)/2

		if arr[mid] == target {
			return mid
		}
		if target > arr[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}
