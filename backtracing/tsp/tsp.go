package tsp

import "algo/graph"
import "fmt"

// g 临接矩阵
// city 除起点外的城市集合
func TSP(g *graph.Graph, s int, city []int, k int, curcost int, best *int, bestpath *[]int) {
	if k == len(city) {
		lastcity := city[len(city)-1]
		if g.Matrix[lastcity][s] != graph.MAXDIST {
			sum := 0
			i := 0
			for ; i < len(city) - 1; i++ {
				sum += g.Matrix[city[i]][city[i+1]]
			}
			sum += g.Matrix[city[i]][s]
			fmt.Println(sum)
			fmt.Println(city)
			if sum < *best {
				*best = sum
				for j := 0; j < len(city); j++ {
					(*bestpath)[j] = city[j]
				}
			}
		}
		return
	}

	for i := k; i < len(city); i++ {
		city[i], city[k] = city[k], city[i]
		curprev, cur := city[k-1], city[k]
		if g.Matrix[curprev][cur] != graph.MAXDIST && curcost + g.Matrix[curprev][cur] < *best {
			TSP(g, s, city, k+1, curcost + g.Matrix[curprev][cur], best, bestpath)
		}
		city[i], city[k] = city[k], city[i]
	}
}