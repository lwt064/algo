package linkedlist

import "fmt"

type Node struct {
	Key interface{}
	Data interface{}
	next *Node
}

type LinkedList struct {
	head *Node
	cur *Node
}

func NewNode(Key interface{}, Data interface{}) *Node {
	return &Node{
		Key: Key,
		Data: Data,
		next: nil,
	}
}

func NewLinkedList() *LinkedList {
	head := NewNode(nil, nil)
	return &LinkedList{head: head, cur: head}
}

func (l *LinkedList) Insert(Key interface{}, Data interface{}) {
	tail := l.head
	for tail.next != nil {
		tail = tail.next
	}
	node := NewNode(Key, Data)
	tail.next = node
}

func (l *LinkedList) Delete(Key interface{}) {
	prev := l.head
	cur := prev.next
	for cur != nil {
		if cur.Key == Key {
			break
		}
		prev = cur
		cur = cur.next
	}

	if cur != nil {
		prev.next = cur.next
	}
}

func (l *LinkedList) Iter() *Node {
	if l.cur != nil {
		l.cur = l.cur.next
	}
	return l.cur
}

func (l *LinkedList) ResetIter() {
	l.cur = l.head
}

func (l *LinkedList) Reverse() {
	head := l.head
	if head == nil || head.next == nil || head.next.next == nil {
		return
	}

	cur := head.next
	for cur != nil && cur.next != nil {
		p := cur.next
		cur.next = cur.next.next
		p.next = head.next
		head.next = p
	}
	return
}

func (l *LinkedList) Range() {
	x := make([]int, 0)
	for p := l.head.next; p != nil; p = p.next {
		x = append(x, p.Data.(int))
	}
	fmt.Println(x)
}
