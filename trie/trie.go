package trie

const MAX_CHAC_NUM = 26

type TrieNode struct {
	c byte
	path int
	child []*TrieNode 
	isEnd bool
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
		c: c,
		path: 0,
		child: make([]*TrieNode, MAX_CHAC_NUM),
		isEnd: false,
	}
	return trieNode
}

func (trie *Trie) Insert(s string) {
	if len(s) == 0 {
		return
	}
	cur := trie.root
	cur.path++
	for j, _ := range(s) {
		idx := int(s[j]-'a')
		if cur.child[idx] == nil {
			trieNode := NewTrieNode(s[j])
			trieNode.path++
			cur.child[idx] = trieNode
			if j == len(s) - 1 {
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
	for j, _ := range(s) {
		idx := int(s[j]-'a')
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
	for j, _ := range(s) {
		idx := int(s[j]-'a')
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
	for j, _ := range(s) {
		idx := int(s[j]-'a')
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
	for _, child := range(node.child) {
		if child != nil {
			getAllWords(child, prefix + string(child.c), result)
		}
	}
}
