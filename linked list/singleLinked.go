package main

import (
	"fmt"
	"strconv"
	"strings"
)

type node struct {
	next  *node
	value int
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) push(value int) {
	newNode := node{nil, value}

	if l.head == nil {
		l.head = &newNode
	} else {
		currentNode := l.head

		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = &newNode
	}
	l.length++

}

func (l *linkedList) display() {
	head := l.head

	var builder strings.Builder

	for head != nil {
		builder.WriteString(strconv.Itoa(head.value))
		builder.WriteString(" -> ")
		head = head.next
	}

	fmt.Println(builder.String())
}

func SingleLinkedList() {
	ll := linkedList{}

	ll.push(10)
	ll.push(11)
	ll.display()

}
