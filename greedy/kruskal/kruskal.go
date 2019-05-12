package kruskal

import (
	"algo/graph"
	"sort"
)

type Edge struct {
	U int // from
	V int // to
	W int // weight
}

type EdgeSet []Edge

func (s EdgeSet) Len() int {
	return len(s)
}
func (s EdgeSet) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s EdgeSet) Less(i, j int) bool {
	return s[i].W < s[j].W
}

// 寻找根节点，如果两个顶点联通，那么他们的根节点相等
func ffind(fa []int, x int) int {
	if fa[x] == x {
		return x
	} else {
		return ffind(fa, fa[x])
	}
}

// 单向图
func Kruskal(g *graph.Graph) EdgeSet {
	es := EdgeSet{}
	for i, _ := range g.Adj {
		cur := g.Adj[i].Iter()
		for cur != nil {
			edge := Edge{i, cur.Key.(int), cur.Data.(int)}
			es = append(es, edge)
			cur = g.Adj[i].Iter()
		}
		g.Adj[i].ResetIter()
	}
	sort.Sort(es)

	result := EdgeSet{}
	fa := make([]int, g.V)
	for i, _ := range fa {
		fa[i] = i
	}
	for _, edge := range es {
		u := edge.U
		v := edge.V
		if ffind(fa, u) != ffind(fa, v) {	// 顶点联通性计算
			result = append(result, edge)
			if len(es) == g.V-1 {
				break
			}
			fa[ffind(fa, v)] = ffind(fa, u)
		}
	}
	return result
}
