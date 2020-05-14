package tree

import (
	"algo/stack"
	"fmt"
)

type TreeNode struct {
	data  interface{}
	left  *TreeNode
	right *TreeNode
}

type Tree struct {
	Root *TreeNode
}

func NewTree() *Tree {
	treeNode := NewTreeNode(nil)
	return &Tree{Root: treeNode}
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

	t.Root = p1
	p1.left = p2
	p1.right = p3
	p2.left = p4
	p3.right = p5
	p4.right = p6
	p6.left = p7
	p6.right = p8

	return t
}

func (t *Tree) PreOrder() []interface{} {
	root := t.Root
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
	root := t.Root
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

func (t *Tree) PostOrder() []interface{} {
	root := t.Root
	if root == nil {
		return nil
	}

	visited := []interface{}{}

	s := stack.NewStack()
	p := root
	lastVisit := p
	for !s.Empty() || p != nil {
		for p != nil {
			s.Push(p)
			p = p.left
		}
		p = s.Top().(*TreeNode)
		if p.right == nil || p.right == lastVisit {
			visited = append(visited, p.data)
			s.Pop()
			lastVisit = p
			p = nil
		} else {
			p = p.right
		}
	}

	return visited
}

func HasSubTree(r1 *TreeNode, r2 *TreeNode) bool {
	if r1 == nil || r2 == nil {
		return false
	}

	result := false
	if r1.data == r2.data {
		result = IsSame(r1, r2)
	}
	if result {
		return true
	}
	return HasSubTree(r1.left, r2) || HasSubTree(r1.right, r2)
}

func IsSame(n1 *TreeNode, n2 *TreeNode) bool {
	if n2 == nil {
		return true
	} else if n1 == nil {
		return false
	}

	if n1.data != n2.data {
		return false
	}
	return IsSame(n1.left, n2.left) && IsSame(n1.right, n2.right)
}

func Mirror(n *TreeNode) *TreeNode {
	if n == nil {
		return nil
	}

	tmp := n.left
	n.left = Mirror(n.right)
	n.right = Mirror(tmp)
	return n
}

func IsSymmetric(t *Tree) bool {
	if t == nil || t.Root == nil {
		return false
	}
	return IsSymmetricCore(t.Root.left, t.Root.right)
}

func IsSymmetricCore(n1 *TreeNode, n2 *TreeNode) bool {
	if n1 == nil && n2 == nil {
		return true
	} else if n1 == nil || n2 == nil {
		return false
	}
	if n1.data != n2.data {
		return false
	}
	return IsSymmetricCore(n1.left, n2.right) && IsSymmetricCore(n1.right, n2.left)
}

func InitSearchTree() *Tree {
	t := NewTree()
	p1 := NewTreeNode(1)
	p2 := NewTreeNode(2)
	p3 := NewTreeNode(3)
	p4 := NewTreeNode(4)
	p5 := NewTreeNode(5)
	p6 := NewTreeNode(6)
	p7 := NewTreeNode(7)
	p8 := NewTreeNode(8)

	t.Root = p4
	p4.left = p2
	p4.right = p5
	p2.left = p1
	p2.right = p3
	p5.right = p7
	p7.left = p6
	p7.right = p8

	return t
}
