package linkedlist

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type LinkedList struct {
	head *Node
}

func NewNode(data interface{}) *Node {
	return &Node{
		data: data,
		next: nil,
	}
}

func InitLinkedList() *LinkedList {
	head := NewNode(nil)
	return &LinkedList{head: head}
}

func (l *LinkedList) Insert(data interface{}) {
	tail := l.head
	for tail.next != nil {
		tail = tail.next
	}
	node := NewNode(data)
	tail.next = node
}

func (l *LinkedList) Delete(data interface{}) {
	prev := l.head
	cur := prev.next
	for cur != nil {
		if cur.data == data {
			break
		}
		prev = cur
		cur = cur.next
	}

	if cur != nil {
		prev.next = cur.next
	}
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
		x = append(x, p.data.(int))
	}
	fmt.Println(x)
}
