package trie

import "algo/queue"

const MAX_CHAC_NUM = 26

type TrieNode struct {
	c      byte
	path   int
	length int
	child  []*TrieNode
	isEnd  bool
	fail   *TrieNode // 增加ac自动机fail指针
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	trieNode := NewTrieNode(0)
	trie := &Trie{root: trieNode}
	return trie
}

func NewTrieNode(c byte) *TrieNode {
	trieNode := &TrieNode{
		c:     c,
		path:  0,
		child: make([]*TrieNode, MAX_CHAC_NUM),
		isEnd: false,
		fail:  nil,
	}
	return trieNode
}

func (trie *Trie) Insert(s string) {
	if len(s) == 0 {
		return
	}
	cur := trie.root
	cur.path++
	for j, _ := range s {
		idx := int(s[j] - 'a')
		if cur.child[idx] == nil {
			trieNode := NewTrieNode(s[j])
			trieNode.path++
			trieNode.length = j + 1
			cur.child[idx] = trieNode
			if j == len(s)-1 {
				trieNode.isEnd = true
				break
			}
		} else {
			cur.child[idx].path++
		}
		cur = cur.child[idx]
	}
	return
}

func (trie *Trie) Delete(s string) {
	if len(s) == 0 {
		return
	}
	found := trie.Find(s)
	if !found {
		return
	}
	cur := trie.root
	cur.path--
	for j, _ := range s {
		idx := int(s[j] - 'a')
		cur.child[idx].path--
		if cur.child[idx].path == 0 {
			cur.child[idx] = nil
			break
		}
		cur = cur.child[idx]
	}
	return
}

func (trie *Trie) Find(s string) bool {
	if len(s) == 0 {
		return false
	}
	cur := trie.root
	for j, _ := range s {
		idx := int(s[j] - 'a')
		if cur.child[idx] == nil {
			return false
		}
		cur = cur.child[idx]
	}
	if !cur.isEnd {
		return false
	}
	return true
}

func (trie *Trie) FindByPrefix(s string) []string {
	if len(s) == 0 {
		return nil
	}
	cur := trie.root
	for j, _ := range s {
		idx := int(s[j] - 'a')
		if cur.child[idx] == nil {
			return nil
		}
		cur = cur.child[idx]
	}

	words := []string{}
	getAllWords(cur, s, &words)
	return words
}

func getAllWords(node *TrieNode, prefix string, result *[]string) {
	if node.isEnd {
		*result = append(*result, prefix)
	}
	for _, child := range node.child {
		if child != nil {
			getAllWords(child, prefix+string(child.c), result)
		}
	}
}

func (trie *Trie) BuildFail() {
	trie.root.fail = nil
	tmpQ := queue.NewLinkedListQueue()
	tmpQ.EnQueue(trie.root)

	for tmpQ.Length() > 0 {
		p := tmpQ.DeQueue().(*TrieNode)
		for i := 0; i < MAX_CHAC_NUM; i++ {
			pc := p.child[i]
			if pc == nil {
				continue
			}

			if p == trie.root {
				pc.fail = p
			} else {
				q := p.fail
				for q != nil {
					idx := int(pc.c - 'a')
					qc := q.child[idx]
					if qc != nil {
						pc.fail = qc
						break
					}
					q = q.fail
				}
				if q == nil {
					pc.fail = trie.root
				}
			}
			tmpQ.EnQueue(pc)
		}
	}
}

func (trie *Trie) Match(text string) []string {
	words := []string{}

	p := trie.root
	for i, _ := range text {
		idx := int(text[i] - 'a')
		for p.child[idx] == nil && p != trie.root {
			p = p.fail
		}
		p = p.child[idx]
		if p == nil {
			p = trie.root
		}
		tmp := p
		for tmp != trie.root {
			if tmp.isEnd {
				words = append(words, text[i-tmp.length+1:i+1])
			}
			tmp = tmp.fail
		}
	}
	return words
}
