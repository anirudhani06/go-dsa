package main

import "fmt"

func main() {

	arr := []int{1, 2, 3}
	swapped := false

	for i := 0; i < len(arr); i++ {
		swapped = false
		for j := 1; j < len(arr)-i; j++ {
			if arr[j] < arr[j-1] {
				temp := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = temp
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}

	fmt.Println(arr)

}
