package main

import "fmt"

func main() {
	arr := []int{8, 5, 11, 22, 33, 1, 0, 4, 5, 9, 88, 55, 1, 0}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quickSort(arr []int, low, high int) {

	if low >= high {
		return
	}

	s := low
	e := high
	m := s + (e-s)/2

	pivot := arr[m]

	for s <= e {
		for arr[s] < pivot {
			s++
		}
		for arr[e] > pivot {
			e--
		}

		if s <= e {
			temp := arr[s]
			arr[s] = arr[e]
			arr[e] = temp
			s++
			e--
		}

	}

	quickSort(arr, low, e)
	quickSort(arr, s, high)

}
