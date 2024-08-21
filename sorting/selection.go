package main

import "fmt"

func main() {

	arr := []int{4, 5, 3, 2, 1}

	for i := 0; i < len(arr); i++ {

		last := len(arr) - i - 1
		maxIndex := getMaxIndex(arr, 0, last)
		swap(arr, maxIndex, last)

	}

	fmt.Println(arr)

}

func getMaxIndex(arr []int, start int, end int) int {

	max := start

	for i := 0; i <= end; i++ {
		if arr[max] < arr[i] {
			max = i
		}

	}
	return max

}

func swap(arr []int, start, end int) {
	temp := arr[start]
	arr[start] = arr[end]
	arr[end] = temp

}
