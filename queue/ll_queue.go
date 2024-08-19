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

type Queue struct {
	head *Node
	tail *Node
	size uint
}

func (q *Queue) EnQueue(data int) {

	newNode := &Node{data: data}
	if q.IsEmpty() {
		q.head = newNode
		q.tail = newNode
	} else {
		tail := q.tail
		q.tail = newNode
		tail.next = q.tail
	}
	q.size++

}

func (q *Queue) DeQueue() (int, error) {

	if q.IsEmpty() {
		return -1, errors.New("queue is empty")
	}

	dequeuedData := q.head.data

	q.head = q.head.next

	if q.head == nil {
		q.tail = nil
	}

	q.size--

	return dequeuedData, nil
}

func (q *Queue) Front() (int, error) {
	if q.IsEmpty() {
		return -1, errors.New("queue is empty")
	}

	return q.tail.data, nil
}

func (q *Queue) Rear() (int, error) {
	if q.IsEmpty() {
		return -1, errors.New("queue is empty")
	}

	return q.head.data, nil
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Display() {
	if q.IsEmpty() {
		return
	}

	var result strings.Builder

	head := q.head

	for head != nil {
		result.WriteString(strconv.Itoa(head.data))
		if head.next != nil {
			result.WriteString(" <- ")
		}
		head = head.next
	}
	fmt.Println(result.String())
}

func main() {

	queue := Queue{}

	queue.EnQueue(10)
	queue.EnQueue(11)
	queue.EnQueue(12)
	queue.EnQueue(13)
	queue.EnQueue(11)
	queue.DeQueue()
	queue.DeQueue()
	queue.Front()
	queue.Rear()

	queue.Display()
}
