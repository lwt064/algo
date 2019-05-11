package colorgraph

import "algo/graph"
import "fmt"

func IsOk(g *graph.Graph, k int, color []int) bool {
	for i := 0; i < g.V; i++ {
		if g.Matrix[i][k] != graph.MAXDIST && color[i] == color[k] {
			return false
		}
	}
	return true
}

func Color(g *graph.Graph, m int, k int, color []int) {
	if k >= g.V {
		fmt.Println(color)
		return
	}
	for i := 1; i <= m; i++ {
		color[k] = i
		if IsOk(g, k, color) {
			Color(g, m, k+1, color)
		}
		color[k] = 0
	}
}