package main

import (
	"fmt"
)

func main() {
	arr := []int{8, 5, 11, 22, 33, 1, 0, 4, 5, 9, 88, 55, 1, 0}
	mergeSort(arr, 0, len(arr))
	fmt.Println(arr)

}

func mergeSort(arr []int, s, e int) {

	if e-s == 1 {
		return
	}

	m := s + (e-s)/2

	mergeSort(arr, s, m)
	mergeSort(arr, m, e)

	mergeInPlace(arr, s, m, e)
}

func mergeInPlace(arr []int, s, m, e int) {
	newArr := make([]int, e-s)

	i := s
	j := m
	k := 0

	for i < m && j < e {
		if arr[i] < arr[j] {
			newArr[k] = arr[i]
			i++

		} else {
			newArr[k] = arr[j]
			j++
		}
		k++
	}

	for i < m {
		newArr[k] = arr[i]
		i++
		k++
	}

	for j < e {
		newArr[k] = arr[j]
		j++
		k++
	}

	for l := 0; l < len(newArr); l++ {
		arr[s+l] = newArr[l]
	}
}
