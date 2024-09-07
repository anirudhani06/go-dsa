package main

import "fmt"

func main() {

	arr := []int{1, 5, 2}

	first := findInMountainArray(2, arr)
	fmt.Print(first)
}

func findInMountainArray(target int, mountainArr []int) int {

	startIndex := 0
	endIndex := len(mountainArr) - 1

	for startIndex < endIndex {
		mid := startIndex + (endIndex-startIndex)/2
		if mountainArr[mid] > mountainArr[mid+1] {
			endIndex = mid
		} else {
			startIndex = mid + 1
		}
	}

	first := search(0, startIndex, mountainArr, target)
	if first != -1 {
		return first
	}
	return search(startIndex, len(mountainArr)-1, mountainArr, target)

}

func search(startIndex int, endIndex int, arr []int, target int) int {

	isAsc := arr[startIndex] < arr[endIndex]

	for startIndex <= endIndex {
		mid := startIndex + (endIndex-startIndex)/2

		if arr[mid] == target {
			return mid
		}

		if isAsc {

			if target < arr[mid] {
				endIndex = mid - 1
			} else {
				startIndex = mid + 1
			}

		} else {
			if target > arr[mid] {
				endIndex = mid - 1
			} else {
				startIndex = mid + 1
			}

		}

	}

	return -1

}
