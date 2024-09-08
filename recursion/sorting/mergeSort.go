package main

import "fmt"

func main() {
	arr := []int{1, 2, 4, 5, 6}

	mergeSort(arr)
	fmt.Println(arr)
}

func mergeSort(arr []int) []int {

	if len(arr) == 1 {
		return arr
	}

	mid := len(arr) / 2

	left := mergeSort(append([]int(nil), arr[0:mid]...))
	right := mergeSort(append([]int(nil), arr[mid:]...))

	return merge(left, right)

}

func merge(left, right []int) []int {
	newArr := make([]int, len(left)+len(right))

	i := 0
	j := 0
	k := 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			newArr[k] = left[i]
			i++
		} else {
			newArr[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		newArr[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		newArr[k] = left[j]
		j++
		k++
	}

	return newArr
}
