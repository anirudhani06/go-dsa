package main

import "fmt"

func main() {
	arr := []int{1, 55, 44, 7, 22, 88, 40}

	bubbleSort(arr, len(arr)-1, 0)
	fmt.Println(arr)
}

func bubbleSort(arr []int, length, c int) {

	if length == 0 {
		return
	}

	if c < length {
		if arr[c] > arr[c+1] {
			temp := arr[c]
			arr[c] = arr[c+1]
			arr[c+1] = temp
		}
		bubbleSort(arr, length, c+1)
	} else {
		bubbleSort(arr, length-1, 0)
	}

}
