package queue

import "fmt"

type ListNode struct {
	Val  interface{}
	next *ListNode
	prev *ListNode
}

type LinkedListQueue struct {
	head   *ListNode // 头哨兵
	tail   *ListNode // 尾哨兵
	length int
}

func NewLinkedListQueue() *LinkedListQueue {
	head := &ListNode{nil, nil, nil}
	tail := &ListNode{nil, nil, nil}
	head.next = tail
	tail.prev = head
	return &LinkedListQueue{head, tail, 0}
}

func (q *LinkedListQueue) EnQueue(v interface{}) {
	node := &ListNode{v, nil, nil}

	node.next = q.head.next
	q.head.next.prev = node
	q.head.next = node
	node.prev = q.head

	q.length++
}

func (q *LinkedListQueue) DeQueue() interface{} {
	if q.head.next == q.tail {
		return nil
	}

	node := q.tail.prev
	prev := node.prev
	prev.next = q.tail
	q.tail.prev = prev

	q.length--

	return node.Val
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
