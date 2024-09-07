package main

import "fmt"

// *
// **
// ***
// ****

func main() {
	pattern(4, 0)
}

func pattern(r, c int) {
	if r == 0 {
		return
	}

	if c < r {
		pattern(r, c+1)
		fmt.Print("*")
	} else {
		pattern(r-1, 0)
		fmt.Println()
	}

}
