package nearpoint

import "math"
import "sort"

type Point struct {
	X float64
	Y float64
}

func Dist(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(float64(p1.X-p2.X), 2) + math.Pow(float64(p1.Y-p2.Y), 2))
}

type PointSetX []Point

func (s PointSetX) Len() int {
	return len(s)
}
func (s PointSetX) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s PointSetX) Less(i, j int) bool {
	return s[i].X < s[j].X
}

type PointSetY []Point

func (s PointSetY) Len() int {
	return len(s)
}
func (s PointSetY) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s PointSetY) Less(i, j int) bool {
	return s[i].Y > s[j].Y
}

func MinDist(points []Point) float64 {
	if len(points) <= 1 || len(points) > 3 {
		return 0
	}
	if len(points) == 2 {
		return Dist(points[0], points[1])
	}
	d1 := Dist(points[0], points[1])
	d2 := Dist(points[0], points[2])
	d3 := Dist(points[1], points[2])

	if d1 <= d2 && d1 <= d3 {
		return d1
	}
	if d2 <= d1 && d2 <= d3 {
		return d2
	}
	if d3 <= d1 && d3 <= d2 {
		return d3
	}

	return 0
}

// 复杂度计算: T(n) = 2(T/2) + O(nlog n), 由于主定理 n**1+e != nlog n，需要用递归树求解，复杂度为 n (log n)**2
func FindNearest(points []Point) float64 {
	if len(points) <= 3 {
		return MinDist(points)
	}

	pointsX := make(PointSetX, 0)
	for _, p := range points {
		pointsX = append(pointsX, p)
	}
	sort.Sort(pointsX)
	x0 := pointsX[(len(pointsX)-1)/2].X
	minL := FindNearest(pointsX[0 : (len(pointsX)+1)/2])
	minR := FindNearest(pointsX[(len(pointsX)+1)/2:])
	min := minL
	if minR < min {
		min = minR
	}

	kPoints := make(PointSetY, 0)
	for _, p := range pointsX {
		if p.X >= x0-min && p.X <= x0+min {
			kPoints = append(kPoints, p)
		}
	}

	k := len(kPoints) // k个候选点
	sort.Sort(kPoints)
	for i := 0; i < k; i++ {
		x := 0
		// 只需与对端6个点比较即可
		for j := i + 1; j < k && j <= i+6+x; j++ {
			if kPoints[j].X < x0 {
				x++
				continue
			}
			if Dist(kPoints[i], kPoints[j]) < min {
				min = Dist(kPoints[i], kPoints[j])
			}
		}
	}

	return min
}
