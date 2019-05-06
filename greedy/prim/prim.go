package prim

import "algo/graph"

const MAXDIST = 2<<31

type Edge struct {
	U int	// from
	V int	// to
	W int	// weight
}

func add(s int, arr *[]int) {
	*arr = append(*arr, s)
}

func remove(s int, arr *[]int) {
	i := -1
	for i, _ = range *arr {
		if (*arr)[i] == s {
			break
		}
	}
	if i >= 0 {
		(*arr)[i] = (*arr)[len(*arr) - 1]
		*arr = (*arr)[0:len(*arr)-1]
	}
}

func IsInSet(s int, arr []int) bool {
	idx := -1
	for i, _ := range arr {
		if arr[i] == s {
			idx = i
			break
		}
	}
	return idx >= 0

}

func Prim(g *graph.Graph) []Edge {
	s := 0   // from 0

	selected := []int{}
	unselected := make([]int, g.V)
	for i, _ := range g.Adj {
		unselected[i] = i
	}
	add(s, &selected)
	remove(s, &unselected)

	edgeSet := []Edge{}
	for {
		if len(unselected) == 0 {
			break
		}
		minDist := MAXDIST
		edge := Edge{-1, -1, -1}
		for _, u := range selected {
			cur := g.Adj[u].Iter()
			for cur != nil {
				v := cur.Key.(int)
				w := cur.Data.(int)
				if IsInSet(v, unselected) {
					if w < minDist {
						minDist = w
						edge.U = u
						edge.V = v
						edge.W = w
					}
				}
				cur = g.Adj[u].Iter()
			}
			g.Adj[u].ResetIter()
		}
		edgeSet = append(edgeSet, edge)
		add(edge.V, &selected)
		remove(edge.V, &unselected)
	}
	return edgeSet
}
