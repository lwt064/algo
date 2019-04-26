package hashMap

import (
	"fmt"
)

const DEFAULT_HASH_MAP_SIZE = 16
const MAX_LOAD_PER_BUCKET = 1

type Node struct {
	key  string
	data interface{}
	prev *Node
	next *Node
}

type HashMap struct {
	bucket []*Node // bucket数组
	size   int     // hashmap元素个数
}

func NewHashMap() *HashMap {
	hash := &HashMap{
		bucket: make([]*Node, DEFAULT_HASH_MAP_SIZE),
		size:   0,
	}
	return hash
}

func HashFunc(key string) uint64 {
	seed := uint64(131)
	code := uint64(0)
	for _, c := range key {
		code = code*seed + uint64(c)
	}
	return code
}

func (hash *HashMap) AddNode(key string, data interface{}) {
	if len(key) == 0 {
		return
	}

	node := hash.FindNode(key)
	if node != nil {
		node.data = data
		return
	}

	hash.resize()

	node = &Node{
		key:  key,
		data: data,
		prev: nil,
		next: nil,
	}
	length := len(hash.bucket)
	idx := HashFunc(node.key) % uint64(length)

	// 初始化头结点
	if hash.bucket[idx] == nil {
		hash.bucket[idx] = &Node{
			key:  "",
			data: nil,
			prev: nil,
			next: nil,
		}
	}

	p := hash.bucket[idx]
	if p.next == nil {
		p.next, node.prev = node, p
	} else {
		for ; p.next != nil; p = p.next {
			if node.key < p.next.key {
				// 插入位置为 p -> node -> p.next
				break
			}
		}

		// 处理最后一个结点
		if p.next == nil {
			node.prev, node.next = p, nil
			p.next = node
		} else {
			// 处理头结点和最后一个结点之外的结点
			node.next, node.prev = p.next, p
			p.next, p.next.prev = node, node
		}
	}
	hash.size = hash.size + 1
}

func (hash *HashMap) FindNode(key string) *Node {
	length := len(hash.bucket)
	idx := HashFunc(key) % uint64(length)
	p := hash.bucket[idx]
	for ; p != nil; p = p.next {
		if p.key == key {
			return p
		}
	}
	return nil
}

func (hash *HashMap) DelNode(key string) {
	node := hash.FindNode(key)
	if node == nil {
		return
	}
	length := len(hash.bucket)
	idx := HashFunc(node.key) % uint64(length)
	p := hash.bucket[idx]
	for ; p != nil; p = p.next {
		if p.key == key {
			if p.next != nil {
				p.next.prev, p.prev.next = p.prev, p.next
			} else {
				p.prev.next = nil
			}
			hash.size = hash.size - 1
		}
	}
}

func (hash *HashMap) Range() {
	for _, p := range hash.bucket {
		if p != nil {
			for q := p.next; q != nil; q = q.next {
				fmt.Println("key: ", q.key, " val: ", q.data)
			}
		}
	}
}

func (hash *HashMap) resize() {
	length := len(hash.bucket)
	if (hash.size+1)/length < MAX_LOAD_PER_BUCKET {
		return
	}
	bucketNew := make([]*Node, 2*length)
	length = len(bucketNew)

	for _, head := range hash.bucket {
		if head != nil {
			for node := head.next; node != nil; node = node.next {
				idx := HashFunc(node.key) % uint64(length)

				// 初始化头结点，头结点不存储数据
				if bucketNew[idx] == nil {
					bucketNew[idx] = &Node{
						key:  "",
						data: nil,
						prev: nil,
						next: nil,
					}
				}

				p := bucketNew[idx]
				if p.next == nil {
					// 当前链表为空
					node.prev, node.next = p, nil
					p.next = node
				} else {
					for ; p.next != nil; p = p.next {
						if node.key < p.next.key {
							// 插入位置为 p -> node -> p.next
							break
						}
					}

					// 处理最后一个结点
					if p.next == nil {
						node.prev, node.next = p, nil
						p.next = node
					} else {
						// 处理头结点和最后一个结点之外的结点
						node.next, node.prev = p.next, p
						p.next, p.next.prev = node, node
					}
				}
			}
		}
	}
	hash.bucket = bucketNew
}
