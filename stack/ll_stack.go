package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	data int
	next *Node
}

type Stack struct {
	top  *Node
	size uint
}

func (s *Stack) push(data int) {
	newNode := &Node{data: data, next: s.top}
	s.top = newNode
	s.size++
}

func (s *Stack) pop() (int, error) {
	if s.isEmpty() {
		return -1, errors.New("stack is empty")
	}

	top := s.top
	s.top = top.next
	s.size--
	return top.data, nil

}

func (s *Stack) peak() (int, error) {

	if s.isEmpty() {
		return -1, errors.New("stack is empty")
	}
	return s.top.data, nil
}

func (s *Stack) isEmpty() bool {
	return s.size == 0
}

func (s *Stack) display() {

	if s.size == 0 {
		return
	}

	var result strings.Builder
	top := s.top

	for top != nil {
		result.WriteString(strconv.Itoa(top.data))

		if top.next != nil {
			result.WriteString(" <- ")
		}
		top = top.next
	}

	fmt.Println(result.String())

}

func main() {
	stack := Stack{}

	stack.push(10)
	stack.push(11)
	stack.peak()

	stack.pop()

	stack.display()

}
