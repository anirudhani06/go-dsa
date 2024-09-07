package main

import "fmt"

func main() {
	str := "malayala"

	isPalindrome := palindrome(str, 0, len(str)-1)
	fmt.Println(isPalindrome)
}

func palindrome(str string, start, end int) bool {

	if start > end {
		return true
	}

	if str[start] != str[end] {
		return false
	}

	return palindrome(str, start+1, end-1)

}
