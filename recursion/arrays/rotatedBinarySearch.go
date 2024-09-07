package main

import "fmt"

func main() {
	arr := []int{5, 6, 7, 8, 9, 1, 2, 3}
	result := search(arr, 9, 0, len(arr)-1)
	fmt.Println(result)

}

func search(arr []int, target, s, e int) int {

	if s > e {
		return -1
	}

	mid := s + (e-s)/2

	if target == arr[mid] {
		return mid
	}

	if arr[s] <= arr[mid] {
		if target >= arr[s] && target <= arr[mid] {
			return search(arr, target, s, mid-1)
		} else {
			return search(arr, target, mid+1, e)
		}
	}

	if target >= arr[mid] && target <= arr[e] {
		return search(arr, target, mid+1, e)
	} else {
		return search(arr, target, s, mid-1)

	}

}
