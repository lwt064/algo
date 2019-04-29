package skiplist

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	//最高层数
	MAX_LEVEL = 16
)

//跳表节点结构体
type skipListNode struct {
	//跳表保存的值
	v interface{}
	//用于排序的分值
	score int
	//层高
	level int
	//每层前进指针
	forwards []*skipListNode
}

//新建跳表节点
func newSkipListNode(v interface{}, score, level int) *skipListNode {
	return &skipListNode{v: v, score: score, forwards: make([]*skipListNode, level, level), level: level}
}

//跳表结构体
type SkipList struct {
	//跳表头结点
	head *skipListNode
	//跳表当前层数
	level int
	//跳表长度
	length int
}

//实例化跳表对象
func NewSkipList() *SkipList {
	//头结点，便于操作
	head := newSkipListNode(0, math.MinInt32, MAX_LEVEL)
	return &SkipList{head, 1, 0}
}

func (sl *SkipList) Insert(v interface{}, score int) {
	if v == nil {
		return
	}

	cur := sl.head
	var update [MAX_LEVEL]*skipListNode

	i := MAX_LEVEL - 1
	for ; i >= 0; i-- {
		for cur.forwards[i] != nil {
			if cur.forwards[i].score == score {
				cur.forwards[i].v = v
				return
			}
			if cur.forwards[i].score > score {
				update[i] = cur
				break
			}
			cur = cur.forwards[i]
		}
		if cur.forwards[i] == nil {
			update[i] = cur
		}
	}

	level := 1
	for i := 1; i < MAX_LEVEL; i++ {
		if rand.Int31()%7 == 1 {
			level++
		}
	}

	newNode := newSkipListNode(v, score, level)
	for i := level - 1; i >= 0; i-- {
		tmp := update[i].forwards[i]
		update[i].forwards[i] = newNode
		newNode.forwards[i] = tmp
	}

	if level > sl.level {
		sl.level = level
	}

	sl.length++
}

func (sl *SkipList) Delete(score int) {
	cur := sl.head
	var update [MAX_LEVEL]*skipListNode
	for i := MAX_LEVEL - 1; i >= 0; i-- {
		update[i] = sl.head
		for cur.forwards[i] != nil {
			if cur.forwards[i].score == score {
				update[i] = cur
				break
			}
			if cur.forwards[i].score > score {
				update[i] = cur
				break
			}
			cur = cur.forwards[i]
		}
	}

	for i := MAX_LEVEL - 1; i >= 0; i-- {
		if update[i].forwards[i] != nil && update[i].forwards[i].score == score {
			update[i].forwards[i] = update[i].forwards[i].forwards[i]
		}
	}

	for i := sl.level - 1; i >= 0; i-- {
		if sl.head.forwards[i] == nil {
			sl.level = i
		} else {
			break
		}
	}

	sl.length--
}

func (sl *SkipList) Find(score int) interface{} {
	cur := sl.head

	i := MAX_LEVEL - 1
	for ; i >= 0; i-- {
		for cur.forwards[i] != nil {
			if cur.forwards[i].score == score {
				return cur.forwards[i].v
			}
			if cur.forwards[i].score > score {
				break
			}
			cur = cur.forwards[i]
		}
	}

	return nil
}

func (sl *SkipList) Range() {
	x := []string{}
	for head := sl.head; head.forwards[0] != nil; head = head.forwards[0] {
		x = append(x, head.forwards[0].v.(string))
	}
	fmt.Println(x)
}
