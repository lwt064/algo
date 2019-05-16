package tree

import "algo/stack"

type TreeNode struct {
	data  interface{}
	left  *TreeNode
	right *TreeNode
}

type Tree struct {
	root *TreeNode
}

func NewTree() *Tree {
	treeNode := NewTreeNode(nil)
	return &Tree{root: treeNode}
}

func NewTreeNode(d interface{}) *TreeNode {
	treeNode := &TreeNode{
		data:  d,
		left:  nil,
		right: nil,
	}
	return treeNode
}

func InitTree() *Tree {
	t := NewTree()
	p1 := NewTreeNode(1)
	p2 := NewTreeNode(2)
	p3 := NewTreeNode(3)
	p4 := NewTreeNode(4)
	p5 := NewTreeNode(5)
	p6 := NewTreeNode(6)
	p7 := NewTreeNode(7)
	p8 := NewTreeNode(8)
	p9 := NewTreeNode(9)

	t.root = p1
	p1.left = p2
	p1.right = p3
	p2.left = p4
	p2.right = p5
	p3.right = p6
	p4.left = p7
	p5.left = p8
	p5.right = p9

	return t
}

func (t *Tree) PreOrder() []interface{} {
	root := t.root
	if root == nil {
		return nil
	}

	visited := []interface{}{}

	s := stack.NewStack()
	p := root
	for !s.Empty() || p != nil {
		for p != nil {
			visited = append(visited, p.data)
			s.Push(p)
			p = p.left
		}
		if !s.Empty() {
			p = s.Top().(*TreeNode)
			s.Pop()
			p = p.right
		}
	}

	return visited
}

func (t *Tree) InOrder() []interface{} {
	root := t.root
	if root == nil {
		return nil
	}

	visited := []interface{}{}

	s := stack.NewStack()
	p := root
	for !s.Empty() || p != nil {
		for p != nil {
			s.Push(p)
			p = p.left
		}
		if !s.Empty() {
			p = s.Top().(*TreeNode)
			visited = append(visited, p.data)
			s.Pop()
			p = p.right
		}
	}

	return visited
}
