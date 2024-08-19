package main

import (
	"fmt"
	"strconv"
	"strings"
)

type DNode struct {
	data int
	next *DNode
	prev *DNode
}

type DList struct {
	head   *DNode
	length int
}

func (l *DList) push(data int) {
	head := l.head

	newNode := &DNode{data: data}
	if head == nil {
		l.head = newNode
		l.length++
		return
	}

	for head != nil {
		if head.next == nil {
			newNode.prev = head
			head.next = newNode
			l.length++
			return
		}
		head = head.next
	}
}

func (l *DList) pop() {
	if l.length == 0 {
		fmt.Println("Empty nodes")
		return
	}

	head := l.head

	for head != nil {
		if head.next == nil {
			head.prev.next = nil
			l.length--
			return
		}
		head = head.next
	}
}

func (l *DList) display() {
	head := l.head
	var next strings.Builder
	var prev strings.Builder

	for head != nil {
		next.WriteString(strconv.Itoa(head.data))

		if head.prev == nil {
			prev.WriteString("nil")
		} else {
			prev.WriteString(" <- ")
			prev.WriteString(strconv.Itoa(head.prev.data))
		}

		if head.next != nil {
			next.WriteString(" -> ")
		}
		head = head.next
	}

	fmt.Println(next.String())
	fmt.Println(prev.String())
}

func doubleLinkedList() {

	dl := DList{}
	dl.push(10)
	dl.push(45)
	dl.push(78)
	dl.pop()

	dl.display()

}
