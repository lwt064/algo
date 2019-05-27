package prim

import "algo/graph"

const MAXDIST = 2<<31

type Edge struct {
	U int	// from
	V int	// to
	W int	// weight
}

func extendV(v int, unselected, selected *[]int) {
	idx := -1
	for i, _ := range *unselected {
		if (*unselected)[i] == v {
			idx = i
			break
		}
	}
	if idx >= 0 {
		(*unselected)[idx] = (*unselected)[len(*unselected)-1]
		*unselected = (*unselected)[0:len(*unselected)-1]
		*selected = append(*selected, v)
	}
}

func IsInSet(s int, unselected []int) bool {
	idx := -1
	for i, _ := range unselected {
		if unselected[i] == s {
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
	extendV(s, &unselected, &selected)

	edgeSet := []Edge{}
	for len(unselected) > 0 {
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
		extendV(edge.V, &unselected, &selected)
		edgeSet = append(edgeSet, edge)
	}
	return edgeSet
}
