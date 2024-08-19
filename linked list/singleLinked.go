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

// O(n) time complexity
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

// O(n) time complexity
func (l *linkedList) pop() {
	head := l.head

	if head == nil {
		return
	}

	if head.next == nil {
		l.head = nil
		l.length--
		return
	}

	for head != nil {
		nextNode := head.next
		if nextNode != nil && nextNode.next == nil {
			head.next = nil

		} else {
			head = head.next
		}

	}
	l.length--

}

// O(1) time complexity
func (l *linkedList) shift(value int) {
	newNode := node{nil, value}

	if l.head == nil {
		l.head = &newNode
	} else {
		newNode.next = l.head
		l.head = &newNode
	}

	l.length++
}

// O(1) time complexity
func (l *linkedList) unshift() {
	head := l.head

	if head == nil {
		return
	}

	if head.next == nil {
		l.head = nil
		l.length--
		return
	}

	l.head = head.next
	l.length--

}

// O(n)
func (l *linkedList) destroy(index int) {
	if index < 0 {
		fmt.Println("Index must be greater than 0.")
		return
	}
	if l.length == 0 {
		fmt.Println("Empty nodes")
		return
	}
	if index >= l.length {
		fmt.Printf("Index %d out of range.\n", index)
		return
	}

	if index == 0 {
		l.head = l.head.next
		return
	}

	i := 1

	head := l.head

	for head != nil {
		if i == index {
			head.next = head.next.next
			l.length--
			break
		}
		i++
		head = head.next
	}
}

// O(n) time complexity
func (l *linkedList) display() {
	head := l.head

	var builder strings.Builder

	for head != nil {
		builder.WriteString(strconv.Itoa(head.value))
		if head.next != nil {
			builder.WriteString(" -> ")
		}
		head = head.next
	}

	fmt.Println(builder.String())
	fmt.Printf("Linked list length: %d\n", l.length)
}

func SingleLinkedList() {
	ll := linkedList{}

	ll.push(10)
	ll.push(11)
	ll.push(60)
	ll.pop()
	ll.push(12)
	ll.unshift()
	ll.shift(100)
	ll.destroy(1)

	ll.display()

}
