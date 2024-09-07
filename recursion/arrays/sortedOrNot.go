package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 400, 55}

	result := isSorted(arr, 0)
	fmt.Println(result)
}

func isSorted(arr []int, i int) bool {

	if i == len(arr)-1 {
		return true
	}

	return arr[i] < arr[i+1] && isSorted(arr, i+1)

}
