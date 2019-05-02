package graph

import "algo/linkedlist"
import "algo/queue"

// 无向带权图的邻接表存储
type Graph struct {
	v int
	adj []*linkedlist.LinkedList
}

func NewGraph(v int) *Graph {
	g := &Graph{
		v: v,
		adj: make([]*linkedlist.LinkedList, v),
	}
	for i := 0; i < g.v; i++ {
		g.adj[i] = linkedlist.NewLinkedList()
	}
	return g
}

func (g *Graph) Insert(s int, to int, weight int) {
	if s >= g.v || to >= g.v {
		return
	}
	g.adj[s].Insert(to, weight)
	g.adj[to].Insert(s, weight)
}

func (g *Graph) Delete(s int, to int) {
	if s >= g.v || to >= g.v {
		return
	}
	g.adj[s].Delete(to)
	g.adj[to].Delete(s)
}

// 广度优先搜索
func (g *Graph) BFS(s int, to int) []int {
	if s == to {
		return nil
	}
	q := queue.NewLinkedListQueue()
	q.EnQueue(s)

	visited := make([]bool, g.v)
	prev := make([]int, g.v)
	for i := 0; i < g.v; i++ {
		visited[i] = false
		prev[i] = -1
	}
	visited[s] = true

	for q.Length() > 0 {
		v := q.DeQueue().(int)
		cur := g.adj[v].Iter()
		for cur != nil {
			adjv := cur.Key.(int)
			if !visited[adjv] {
				prev[adjv] = v
				if adjv == to {
					g.adj[v].ResetIter()
					return prev
				}
				visited[adjv] = true
				q.EnQueue(adjv)
			}
			cur = g.adj[v].Iter()
		}
		g.adj[v].ResetIter()
	}
	return nil
}

func (g *Graph) DFS(s int, to int) []int {
	visited := make([]bool, g.v)
	prev := make([]int, g.v)
	for i := 0; i < g.v; i++ {
		visited[i] = false
		prev[i] = -1
	}
	visited[s] = true

	found := false
	g.recurDFS(s, to, visited, &prev, &found)
	return prev
}

func (g *Graph) recurDFS(s int, to int, visited []bool, prev *[]int, found *bool) {
	if *found {
		return
	}

	visited[s] = true
	if s == to {
		*found = true
		return
	}

	// 遍历邻接表
	cur := g.adj[s].Iter()
	for cur != nil {
		adjv := cur.Key.(int)
		if !*found && !visited[adjv] {
			(*prev)[adjv] = s
			g.recurDFS(adjv, to, visited, prev, found)
		}
		cur = g.adj[s].Iter()
	}
	g.adj[s].ResetIter()
}

// 查找n度以内的好友，包含n度
func (g *Graph) FindNFriends(s int, n int) []int {
	q := queue.NewLinkedListQueue()
	q.EnQueue(s)

	visited := make([]bool, g.v)
	dist := make([]int, g.v)
	for i := 0; i < g.v; i++ {
		visited[i] = false
		dist[i] = -1
	}
	visited[s] = true
	dist[s] = 0

	friends := []int{}
	for q.Length() > 0 {
		v := q.DeQueue().(int)
		if dist[v] > n {
			break
		}
		cur := g.adj[v].Iter()
		for cur != nil {
			adjv := cur.Key.(int)
			if !visited[adjv] {
				dist[adjv] = dist[v] + 1
				if dist[adjv] <= n {
					friends = append(friends, adjv)
				}
				visited[adjv] = true
				q.EnQueue(adjv)
			}
			cur = g.adj[v].Iter()
		}
		g.adj[v].ResetIter()
	}
	return friends
}