package linkedlist

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

func InitLinkedList() *Node {
	head := &Node{
		data: nil,
		next: nil,
	}
	cur := head
	for i := 0; i < 10; i++ {
		node := &Node{
			data: i,
			next: nil,
		}
		cur.next = node
		cur = cur.next
	}
	return head
}

func Reverse(head *Node) *Node {
	if head == nil || head.next == nil || head.next.next == nil {
		return head
	}

	cur := head.next
	for cur != nil && cur.next != nil {
		p := cur.next
		cur.next = cur.next.next
		p.next = head.next
		head.next = p
	}

	return head
}

func Range(head *Node) {
	x := make([]int, 0)
	for p := head.next; p != nil; p = p.next {
		x = append(x, p.data.(int))
	}
	fmt.Println(x)
}
