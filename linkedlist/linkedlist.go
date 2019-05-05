package linkedlist

import "fmt"

type Node struct {
	Key  interface{}
	Data interface{}
	Next *Node
}

type LinkedList struct {
	head *Node
	cur  *Node
}

func NewNode(Key interface{}, Data interface{}) *Node {
	return &Node{
		Key:  Key,
		Data: Data,
		Next: nil,
	}
}

func NewLinkedList() *LinkedList {
	head := NewNode(nil, nil)
	return &LinkedList{head: head, cur: head}
}

func (l *LinkedList) Insert(Key interface{}, Data interface{}) {
	tail := l.head
	for tail.Next != nil {
		tail = tail.Next
	}
	node := NewNode(Key, Data)
	tail.Next = node
}

func (l *LinkedList) Delete(Key interface{}) {
	prev := l.head
	cur := prev.Next
	for cur != nil {
		if cur.Key == Key {
			break
		}
		prev = cur
		cur = cur.Next
	}

	if cur != nil {
		prev.Next = cur.Next
	}
}

func (l *LinkedList) Iter() *Node {
	if l.cur != nil {
		l.cur = l.cur.Next
	}
	return l.cur
}

func (l *LinkedList) ResetIter() {
	l.cur = l.head
}

func (l *LinkedList) Reverse() {
	head := l.head
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}

	cur := head.Next
	for cur != nil && cur.Next != nil {
		p := cur.Next
		cur.Next = cur.Next.Next
		p.Next = head.Next
		head.Next = p
	}
	return
}

func (l *LinkedList) Range() {
	x := make([]int, 0)
	for p := l.head.Next; p != nil; p = p.Next {
		x = append(x, p.Data.(int))
	}
	fmt.Println(x)
}
