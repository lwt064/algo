package maze

import "algo/queue"
import "fmt"

// 迷宫问题
type Maze struct {
	pMap [][]int
	xSzie int
	ySzie int
}

type Point struct {
	X int
	Y int
}

func NewMaze() *Maze {
	pMap := [][]int{
		{0, 1, 0, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 0, 0, 0, 0},
		{0, 1, 1, 1, 0},
		{0, 0, 0, 1, 0},
	}
	return &Maze{
		pMap: pMap,
		xSzie: 5,
		ySzie: 5,
	}
}

func (maze *Maze) canMove(x int, y int) bool {
	if x < 0 || x >= maze.xSzie {
		return false
	}
	if y < 0 || y >= maze.ySzie {
		return false
	}
	if maze.pMap[x][y] == 1 {
		return false
	}
	return true
}

func (maze *Maze) BFS(from, to Point) [][]Point {
	q := queue.NewLinkedListQueue()
	q.EnQueue(from)

	visited := make([][]bool, maze.xSzie)
	prev := make([][]Point, maze.xSzie)
	for i := 0; i < maze.xSzie; i++ {
		visited[i] = make([]bool, maze.ySzie)
		prev[i] = make([]Point, maze.ySzie)

		for j := 0; j < maze.ySzie; j++ {
			visited[i][j] = false
			prev[i][j] = Point{-1, -1}
		}
	}
	visited[from.X][from.Y] = true

	dist := []Point{
		Point{1, 0},	// 往下
		Point{0, 1},    // 往右
		Point{-1, 0},   // 往上
		Point{0, -1},   // 往左
	}

	for q.Length() > 0 {
		cur := q.DeQueue().(Point)
		for i := 0; i < len(dist); i++ {
			next := Point{
				cur.X + dist[i].X,
				cur.Y + dist[i].Y,
			}
			if maze.canMove(next.X, next.Y) && !visited[next.X][next.Y] {
				prev[next.X][next.Y] = Point{cur.X, cur.Y}
				visited[next.X][next.Y] = true
				if next.X == to.X && next.Y == to.Y {
					return prev
				}
				q.EnQueue(next)
			}	
		}
	}
	return nil
}

func (maze *Maze) PrintPath(prevPath [][]Point, from, to Point) {
	positivePath := []Point{to}
	prev := prevPath[to.X][to.Y]
	for prev.X != -1 {
		if prev.X == from.X && prev.Y == from.Y {
			break
		}
		positivePath = append(positivePath, prev)
		prev = prevPath[prev.X][prev.Y]
	}
	nPath := len(positivePath)
	for i := 0; i < nPath/2 ; i++ {
		positivePath[i], positivePath[nPath-i-1] = positivePath[nPath-i-1], positivePath[i]
	}
	fmt.Println(positivePath)
}