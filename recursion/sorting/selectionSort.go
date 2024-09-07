package main

import "fmt"

func main() {
	arr := []int{1, 77, 55, 10, 21, 36, 41}

	selectionSort(arr, len(arr), 0, 0)
	fmt.Println(arr)
}

func selectionSort(arr []int, length, i, max int) {
	if length == 0 {
		return
	}

	if i < length {
		if arr[i] > arr[max] {
			selectionSort(arr, length, i+1, i)

		} else {
			selectionSort(arr, length, i+1, max)
		}
	} else {
		swap(arr, max, length-1)
		selectionSort(arr, length-1, 0, 0)
	}
}

func swap(arr []int, s, e int) {
	temp := arr[s]
	arr[s] = arr[e]
	arr[e] = temp
}
