package queue

import "fmt"

type ListNode struct {
	Val  interface{}
	next *ListNode
}

type LinkedListQueue struct {
	head   *ListNode
	length int
}

func NewLinkedListQueue() *LinkedListQueue {
	head := &ListNode{nil, nil}
	return &LinkedListQueue{head, 0}
}

func (q *LinkedListQueue) EnQueue(v interface{}) {
	node := &ListNode{v, nil}
	node.next = q.head.next
	q.head.next = node
	q.length++
}

func (q *LinkedListQueue) DeQueue() interface{} {
	if q.head == nil || q.head.next == nil {
		return nil
	}
	prev := q.head
	tail := prev.next
	for tail != nil && tail.next != nil {
		prev = tail
		tail = tail.next
	}

	prev.next = nil
	q.length--
	return tail.Val
}

func (q *LinkedListQueue) Length() int {
	return q.length
}

func (q *LinkedListQueue) Range() {
	if q.head == nil {
		return
	}
	x := []interface{}{}
	for cur := q.head; cur != nil; cur = cur.next {
		x = append(x, cur.Val)
	}
	fmt.Println(x)
}