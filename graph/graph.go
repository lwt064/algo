package graph

import "algo/linkedlist"
import "algo/queue"
import "fmt"

const DIRECTION_SINGLE = 1
const DIRECTION_BOTH = 2

const MAXDIST = 99999999

// 单向/双向带权图的邻接表存储
type Graph struct {
	direction int
	V         int
	Adj       []*linkedlist.LinkedList		// 邻接表
	Matrix    [][]int						// 临接矩阵
}

func NewGraph(V int, direction int) *Graph {
	g := &Graph{
		direction: direction,
		V:         V,
		Adj:       make([]*linkedlist.LinkedList, V),
		Matrix:    make([][]int, V),
	}
	for i := 0; i < g.V; i++ {
		g.Adj[i] = linkedlist.NewLinkedList()
		g.Matrix[i] = make([]int, V)
		for j, _ := range g.Matrix[i] {
			g.Matrix[i][j] = MAXDIST
		}
	}
	return g
}

func (g *Graph) Insert(s int, to int, weight int) {
	if s >= g.V || to >= g.V {
		return
	}
	if g.direction == DIRECTION_BOTH {
		g.Adj[s].Insert(to, weight)
		g.Adj[to].Insert(s, weight)
		g.Matrix[s][to] = weight
		g.Matrix[to][s] = weight
	} else {
		g.Adj[s].Insert(to, weight)
		g.Matrix[s][to] = weight
	}
}

func (g *Graph) Delete(s int, to int) {
	if s >= g.V || to >= g.V {
		return
	}
	if g.direction == DIRECTION_BOTH {
		g.Adj[s].Delete(to)
		g.Adj[to].Delete(s)
		g.Matrix[s][to] = MAXDIST
		g.Matrix[to][s] = MAXDIST
	} else {
		g.Adj[s].Delete(to)
		g.Matrix[s][to] = MAXDIST
	}
}

// 广度优先搜索
func (g *Graph) BFS(s int, to int) []int {
	if s == to {
		return nil
	}
	q := queue.NewLinkedListQueue()
	q.EnQueue(s)

	visited := make([]bool, g.V)
	prev := make([]int, g.V)
	for i := 0; i < g.V; i++ {
		visited[i] = false
		prev[i] = -1
	}
	visited[s] = true

	for q.Length() > 0 {
		V := q.DeQueue().(int)
		cur := g.Adj[V].Iter()
		for cur != nil {
			adjv := cur.Key.(int)
			if !visited[adjv] {
				prev[adjv] = V
				if adjv == to {
					g.Adj[V].ResetIter()
					return prev
				}
				visited[adjv] = true
				q.EnQueue(adjv)
			}
			cur = g.Adj[V].Iter()
		}
		g.Adj[V].ResetIter()
	}
	return nil
}

func (g *Graph) DFS(s int, to int) []int {
	visited := make([]bool, g.V)
	prev := make([]int, g.V)
	for i := 0; i < g.V; i++ {
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
	cur := g.Adj[s].Iter()
	for cur != nil {
		adjv := cur.Key.(int)
		if !*found && !visited[adjv] {
			(*prev)[adjv] = s
			g.recurDFS(adjv, to, visited, prev, found)
		}
		cur = g.Adj[s].Iter()
	}
	g.Adj[s].ResetIter()
}

// 查找n度以内的好友，包含n度
func (g *Graph) FindNFriends(s int, n int) []int {
	q := queue.NewLinkedListQueue()
	q.EnQueue(s)

	visited := make([]bool, g.V)
	dist := make([]int, g.V)
	for i := 0; i < g.V; i++ {
		visited[i] = false
		dist[i] = -1
	}
	visited[s] = true
	dist[s] = 0

	friends := []int{}
	for q.Length() > 0 {
		V := q.DeQueue().(int)
		if dist[V] > n {
			break
		}
		cur := g.Adj[V].Iter()
		for cur != nil {
			adjv := cur.Key.(int)
			if !visited[adjv] {
				dist[adjv] = dist[V] + 1
				if dist[adjv] <= n {
					friends = append(friends, adjv)
				}
				visited[adjv] = true
				q.EnQueue(adjv)
			}
			cur = g.Adj[V].Iter()
		}
		g.Adj[V].ResetIter()
	}
	return friends
}

func (g *Graph) PrintPath(prevPath []int, to int) {
	positivePath := []int{to}
	prev := prevPath[to]
	for prev != -1 {
		positivePath = append(positivePath, prev)
		prev = prevPath[prev]
	}
	nPath := len(positivePath)
	for i := 0; i < nPath/2; i++ {
		positivePath[i], positivePath[nPath-i-1] = positivePath[nPath-i-1], positivePath[i]
	}
	fmt.Println(positivePath)
}
