package main

import "fmt"

func main() {

	arr := []int{1, 5, 22, 44, 8, 1, 3, 6, 8, 5, 9}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
			}
		}
	}
	fmt.Println(arr)

}
