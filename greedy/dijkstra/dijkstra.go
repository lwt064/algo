package dijkstra

import "algo/graph"

const MAXDIST = 2 << 31

func MinDist(g *graph.Graph, from int) ([]int, []int) {
	dist, prev, unselectedSet := []int{}, []int{}, []int{}
	for i, _ := range g.Adj {
		dist = append(dist, MAXDIST)
		prev = append(prev, -1)
		unselectedSet = append(unselectedSet, i)
	}
	dist[from] = 0

	for len(unselectedSet) > 0 {
		min := MAXDIST
		idx := 0
		for i, x := range unselectedSet {
			if dist[x] < min {
				min = dist[x]
				idx = i
			}
		}

		v := unselectedSet[idx]
		unselectedSet[idx] = unselectedSet[len(unselectedSet)-1]
		unselectedSet = unselectedSet[0 : len(unselectedSet)-1]

		vv := g.Adj[v].Iter()
		for vv != nil {
			if dist[v]+vv.Data.(int) < dist[vv.Key.(int)] {
				dist[vv.Key.(int)] = dist[v] + vv.Data.(int)
				prev[vv.Key.(int)] = v
			}
			vv = g.Adj[v].Iter()
		}
		g.Adj[v].ResetIter()
	}
	return dist, prev
}
