package main

import "fmt"

// ****
// ***
// **
// *

func main() {
	pattern(4, 0)
}

func pattern(r, c int) {
	if r == 0 {
		return
	}
	if c < r {
		fmt.Print("*")
		pattern(r, c+1)
	} else {
		fmt.Println()
		pattern(r-1, 0)
	}

}
