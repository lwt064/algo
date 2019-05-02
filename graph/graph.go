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
	prev := []int{s}

	visited := make([]bool, g.v)
	for i := 0; i < g.v; i++ {
		visited[i] = false
	}
	visited[s] = true

	for q.Length() > 0 {
		v := q.DeQueue().(int)
		cur := g.adj[v].Iter()
		for cur != nil {
			adjv := cur.Key.(int)
			if !visited[adjv] {
				prev = append(prev, adjv)
				if adjv == to {
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