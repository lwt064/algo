package bin_search_tree

type Node struct {
	data  interface{}
	score int
	left  *Node
	right *Node
}

type BST struct {
	Root *Node
}

func NewBST() *BST {
	return &BST{
		Root: nil,
	}
}

func NewNode(data interface{}, score int) *Node {
	return &Node{
		data:  data,
		score: score,
		left:  nil,
		right: nil,
	}
}

func (bst *BST) Find(score int) interface{} {
	Root := bst.Root
	for Root != nil {
		if Root.score == score {
			return Root.data
		} else if Root.score > score {
			Root = Root.left
		} else {
			Root = Root.right
		}
	}
	return nil
}

func (bst *BST) Insert(data interface{}, score int) {
	if data == nil {
		return
	}
	if bst.Root == nil {
		bst.Root = NewNode(data, score)
		return
	}

	Root := bst.Root
	for Root != nil {
		if Root.score == score {
			Root.data = data
			break
		} else if Root.score > score {
			if Root.left == nil {
				Root.left = NewNode(data, score)
				break
			}
			Root = Root.left
		} else {
			if Root.right == nil {
				Root.right = NewNode(data, score)
				break
			}
			Root = Root.right
		}
	}
	return
}

func (bst *BST) Delete(score int) {
	if bst.Root == nil {
		return
	}

	node := bst.Root // 待删除结点
	var pNode *Node  // 待删除结点的父节点
	for node != nil {
		if node.score == score {
			break
		} else if node.score > score {
			pNode = node
			node = node.left
		} else {
			pNode = node
			node = node.right
		}
	}

	// 没找到
	if node == nil {
		return
	}

	// 左右儿子均为空
	if node.left == nil && node.right == nil {
		if node == bst.Root {
			bst.Root = nil
		} else if pNode.left == node {
			pNode.left = nil
		} else {
			pNode.right = nil
		}
		return
	} else if node.right == nil {
		if node == bst.Root {
			bst.Root = bst.Root.left
		} else if pNode.left == node {
			pNode.left = node.left
		} else {
			pNode.right = node.left
		}
		return
	} else if node.left == nil {
		if node == bst.Root {
			bst.Root = bst.Root.right
		} else if pNode.left == node {
			pNode.left = node.right
		} else {
			pNode.right = node.right
		}
		return
	} else {
		leftMax := node.left
		for leftMax.right != nil {
			pNode = leftMax
			leftMax = leftMax.right
		}
		node.data = leftMax.data
		node.score = leftMax.score
		pNode.right = nil
	}
}

func (bst *BST) Min() interface{} {
	Root := bst.Root
	if Root == nil {
		return nil
	}

	data := Root.data
	for Root != nil {
		if Root.left != nil {
			data = Root.left.data
		}
		Root = Root.left
	}
	return data
}

func (bst *BST) Max() interface{} {
	Root := bst.Root
	if Root == nil {
		return nil
	}

	data := Root.data
	for Root != nil {
		if Root.right != nil {
			data = Root.right.data
		}
		Root = Root.right
	}
	return data
}

func PreOrder(Root *Node, x *[]string) {
	if Root == nil {
		return
	}
	PreOrder(Root.left, x)
	*x = append(*x, Root.data.(string))
	PreOrder(Root.right, x)
	return
}
