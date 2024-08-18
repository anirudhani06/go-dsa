package main

import "fmt"

type Stack struct {
	data   []int
	length int
}

func (s *Stack) push(val int) int {
	s.data = append(s.data, val)
	s.length++
	return val
}

func (s *Stack) pop() int {

	if len(s.data) == 0 {
		return -1
	}

	lastIndex := len(s.data) - 1
	lastItem := s.data[lastIndex]
	s.data = s.data[:lastIndex]
	s.length--
	return lastItem
}

func (s *Stack) peek() int {
	peakElem := s.data[s.length-1]
	return peakElem

}

func (s *Stack) isEmpty() bool {
	return s.length == 0
}

func (s *Stack) size() int {
	return s.length
}

func (s *Stack) show() []int {
	return s.data
}

func main() {
	stack := Stack{}

	stack.push(10)
	stack.push(12)
	stack.pop()
	stack.push(13)

	output := stack.show()

	fmt.Println(output)
}
