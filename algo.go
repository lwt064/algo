package main

import (
	"algo/backtracing/bag01"
	"algo/backtracing/nqueen"

	// "algo/backtracing/arrange"
	"algo/backtracing/colorgraph"
	"algo/backtracing/editdist"
	"algo/backtracing/tsp"
	"algo/bin_search"
	"algo/bin_search_tree"
	"algo/cmpsort"
	"algo/devideconquer/bign"
	"algo/devideconquer/nearpoint"
	"algo/devideconquer/selectk"
	"algo/dynamicprogramming/arraymaxsum"
	"algo/dynamicprogramming/coin"
	"algo/dynamicprogramming/dpbag01"
	"algo/dynamicprogramming/dpeditdist"
	"algo/dynamicprogramming/lcs"
	"algo/dynamicprogramming/lis"
	"algo/dynamicprogramming/yanghui"
	"algo/graph"
	"algo/graph/maze"
	"algo/greedy/dijkstra"
	"algo/greedy/kruskal"
	"algo/greedy/prim"
	hashMap "algo/hashmap"
	"algo/linkedlist"
	"algo/recursive/cellsplit"
	"algo/skiplist"
	"algo/stringmatch"
	"algo/tree"
	"algo/trie"
	"fmt"
	"math/rand"
	"sort"
)

func sortTest() {
	x := make([]int, 0)
	n := 100
	for i := 0; i < n; i++ {
		x = append(x, rand.Intn(n))
	}
	// x = cmpsort.MergeSort(x)
	// // cmpsort.HeapSort(x)
	cmpsort.QuicSort(x)

	fmt.Println("kth: ", cmpsort.SelectK(x, 18))

	// cmpsort.InsertSort(x)
	for _, v := range x {
		fmt.Println(v)
	}
}

func binSearchTest() {
	x := make([]int, 0)
	n := 1000
	for i := 0; i < n; i++ {
		x = append(x, rand.Intn(n))
	}
	sort.Ints(x)
	fmt.Println(bin_search.Bin_search(x, 1))
	fmt.Println(bin_search.Bin_search_First_Equal(x, 10))
	fmt.Println(bin_search.Bin_search_First_GT_Key(x, 1))
}

func hashMapTest() {
	fmt.Println("hash code: ", hashMap.HashFunc("winter"))

	hmap := hashMap.NewHashMap()
	hmap.AddNode("winter", 18)
	hmap.AddNode("comming", 31)
	hmap.AddNode("Is", 20)
	hmap.AddNode("winter", 9)

	fmt.Println("\nAfter Add: ")
	hmap.Range()
	hmap.DelNode("Is")

	fmt.Println("\nAfter Del: ")
	hmap.Range()

	hmap.AddNode("winter1", 18)
	hmap.AddNode("comming1", 31)
	hmap.AddNode("Is1", 20)
	hmap.AddNode("winter2", 18)
	hmap.AddNode("comming2", 31)
	hmap.AddNode("Is2", 20)
	hmap.AddNode("winter3", 18)
	hmap.AddNode("comming3", 31)
	hmap.AddNode("Is3", 20)
	hmap.AddNode("winte4", 18)
	hmap.AddNode("comming4", 31)
	hmap.AddNode("Is4", 20)
	hmap.AddNode("winter5", 9)
	hmap.AddNode("comming5", 31)
	hmap.AddNode("Is5", 20)
	hmap.AddNode("winter6", 9)
	hmap.AddNode("comming6", 31)
	hmap.AddNode("Is6", 20)

	fmt.Println("\nAfter Resize: ")
	hmap.Range()
}

func linkedlistTest() {
	l := linkedlist.NewLinkedList()
	l.Insert(10, 10)
	l.Insert(9, 9)
	l.Insert(8, 8)
	l.Insert(7, 7)
	l.Insert(6, 6)
	l.Insert(5, 5)
	l.Insert(4, 4)
	l.Insert(3, 3)
	l.Insert(2, 2)
	l.Insert(1, 1)
	l.Insert(0, 0)

	l.Range()
	l.Reverse()
	l.Range()

	l.Delete(10)
	l.Delete(5)
	l.Delete(0)

	l.Range()
}

func skipListTest() {
	sl := skiplist.NewSkipList()
	sl.Insert("Winter", 100)
	sl.Insert("Is", 20)
	sl.Insert("Comming", 50)
	sl.Insert("Soon", 30)

	sl.Insert("Winter1", 90)
	sl.Insert("Is1", 25)
	sl.Insert("Comming1", 55)
	sl.Insert("Soon1", 35)

	sl.Insert("Winter2", 70)
	sl.Insert("Is2", 40)
	sl.Insert("Comming2", 60)
	sl.Insert("Soon2", 65)

	sl.Range()

	sl.Delete(25)
	sl.Delete(30)
	sl.Delete(40)

	sl.Range()

	sl.Delete(20)
	sl.Delete(90)
	sl.Delete(100)

	sl.Range()
}

func binSarchTreeTest() {
	bst := bin_search_tree.NewBST()
	bst.Insert("Winter", 10)
	bst.Insert("Is", 20)
	bst.Insert("Comming", 30)
	bst.Insert("Soon", 40)

	bst.Insert("Winter1", 11)
	bst.Insert("Is1", 21)
	bst.Insert("Comming1", 31)
	bst.Insert("Soon1", 41)

	bst.Insert("Winter2", 12)
	bst.Insert("Is2", 22)
	bst.Insert("Comming2", 32)
	bst.Insert("Soon2", 42)

	bst.Insert("Winter3", 13)
	bst.Insert("Is3", 23)
	bst.Insert("Comming3", 33)
	bst.Insert("Soon3", 43)

	x := make([]string, 0)
	bin_search_tree.PreOrder(bst.Root, &x)
	fmt.Println(x)

	fmt.Println("Find: ", bst.Find(41))
	fmt.Println("Min: ", bst.Min())
	fmt.Println("Max: ", bst.Max())

	bst.Delete(42)
	bst.Delete(30)
	bst.Delete(12)

	x = x[:0]
	bin_search_tree.PreOrder(bst.Root, &x)
	fmt.Println(x)
}

func cellSplitTest() {
	fmt.Println(cellsplit.Split1(4))
	fmt.Println(cellsplit.Split2(4))
}

func treeTest() {
	t := tree.InitTree()
	fmt.Println(t.PreOrder())

	fmt.Println(t.InOrder())
}

func graphTest() {
	g := graph.NewGraph(10, graph.DIRECTION_BOTH)
	g.Insert(0, 1, 1)
	g.Insert(0, 3, 3)
	g.Insert(1, 4, 5)
	g.Insert(1, 5, 6)
	g.Insert(1, 6, 7)
	g.Insert(2, 3, 4)
	g.Insert(3, 1, 3)
	g.Insert(3, 6, 6)
	g.Insert(5, 7, 7)
	g.Insert(6, 8, 9)
	g.Insert(7, 8, 7)
	g.Insert(8, 9, 8)

	s := 1
	to := 9

	{
		prevPath := g.BFS(s, to)
		g.PrintPath(prevPath, to)
	}

	{
		prevPath := g.DFSSearch(s, to)
		g.PrintPath(prevPath, to)
	}

	{
		prevPath := g.DFSTravel(s)
		fmt.Println(prevPath)
	}

	{
		friends := g.FindNFriends(s, 2)
		fmt.Println(friends)
	}

	{
		pMaze := maze.NewMaze()
		prevPath := pMaze.BFS(maze.Point{0, 0}, maze.Point{4, 4})
		pMaze.PrintPath(prevPath, maze.Point{0, 0}, maze.Point{4, 4})
	}
}

func stringMatchTest() {
	{
		s := "wintercomming"
		p := "rcom"
		idx := stringmatch.RKMatch(s, p)
		fmt.Println(idx)
	}

	{
		// p := "cabcab"
		// bc := stringmatch.GenerateBC(p)
		// fmt.Println(bc)

		// suffix, prefix := stringmatch.GenerateGS(p)
		// fmt.Println(suffix)
		// fmt.Println(prefix)
	}

	{
		s := "I ove love you, bless China, blessbless my home."
		p := "blessbless"
		fmt.Println(stringmatch.BMMatch(s, p))
	}

	{
		p := "ababaca"
		fmt.Println(stringmatch.GetNext(p))

		s := "ababababadadababacabless"
		fmt.Println(stringmatch.KMPMatch(s, p))
	}
}

func trieTest() {
	{
		trie := trie.NewTrie()
		trie.Insert("hello")
		trie.Insert("hi")
		trie.Insert("how")
		trie.Insert("see")
		trie.Insert("so")

		fmt.Println(trie.Find("see"))
		fmt.Println(trie.Find("hello"))
		fmt.Println(trie.Find("say"))

		fmt.Println(trie.FindByPrefix("h"))

		trie.Delete("hi")

		fmt.Println(trie.FindByPrefix("h"))
	}

	{
		trie := trie.NewTrie()
		trie.Insert("abcd")
		trie.Insert("bcd")
		trie.Insert("c")
		trie.BuildFail()
		fmt.Println(trie.Match("bcdcabcd"))
	}
}

func greedyTest() {
	g := graph.NewGraph(10, graph.DIRECTION_BOTH)
	g.Insert(0, 1, 1)
	g.Insert(0, 3, 3)
	g.Insert(1, 4, 5)
	g.Insert(1, 5, 6)
	g.Insert(1, 6, 7)
	g.Insert(2, 3, 4)
	g.Insert(3, 1, 3)
	g.Insert(3, 6, 6)
	g.Insert(5, 7, 7)
	g.Insert(6, 8, 9)
	g.Insert(7, 8, 7)
	g.Insert(8, 9, 8)

	{
		dist, prev := dijkstra.MinDist(g, 0)
		fmt.Println(dist)
		fmt.Println(prev)
	}

	{
		edgeSet := prim.Prim(g)
		fmt.Println(edgeSet)
	}

	{
		g := graph.NewGraph(10, graph.DIRECTION_SINGLE)
		g.Insert(0, 1, 1)
		g.Insert(0, 3, 3)
		g.Insert(1, 4, 5)
		g.Insert(1, 5, 6)
		g.Insert(1, 6, 7)
		g.Insert(2, 3, 4)
		g.Insert(3, 1, 3)
		g.Insert(3, 6, 6)
		g.Insert(5, 7, 7)
		g.Insert(6, 8, 9)
		g.Insert(7, 8, 7)
		g.Insert(8, 9, 8)
		edgeSet := kruskal.Kruskal(g)
		fmt.Println(edgeSet)
	}
}

// devide and conquer test
func dcTest() {
	points := []nearpoint.Point{
		nearpoint.Point{9.83, -81.96},
		nearpoint.Point{-88.29, 44.76},
		nearpoint.Point{21.97, -81.49},
		nearpoint.Point{2.44, -1.83},
		nearpoint.Point{-89.17, 63.58},
		nearpoint.Point{20, -49.92},
		nearpoint.Point{-81.21, -48.01},
		nearpoint.Point{-33.28, -49.09},
		nearpoint.Point{-54.05, 12.88},
		nearpoint.Point{-64.85, -53.12},
		nearpoint.Point{12.07, 64.91},
		nearpoint.Point{-72.9, -21.57},
		nearpoint.Point{12.93, -92.71},
		nearpoint.Point{-27.71, -0.19},
		nearpoint.Point{73.17, 31.17},
	}
	dist := nearpoint.FindNearest(points)
	fmt.Println(dist)

	{
		x := bign.BigN{1, "9527"}
		y := bign.BigN{0, "9528"}
		z := bign.SmallMultiply(x, y)
		fmt.Println(z.Data)
	}

	{
		x := bign.BigN{1, "987654321"}
		y := bign.BigN{0, "12345678"}
		z := bign.Add(x, y)
		fmt.Println(z.Data)
	}

	{
		x := bign.BigN{0, "12345678900"}
		y := bign.BigN{0, "98765432100"}
		z := bign.Multiply(x, y)
		fmt.Println(z.Data)
	}

	{
		num := 1000
		A := make([]int, num)
		for i := 0; i < num; i++ {
			A[i] = rand.Intn(num)
		}

		k := 46
		fmt.Println(selectk.SelectK(A, k))

		sort.Ints(A)
		fmt.Println(A[len(A)-k])
	}
}

func btTest() {
	{
		fmt.Println(nqueen.Resolve(8))
	}

	{
		w := 25
		b := []int{3, 4, 9, 12, 6, 3}
		fmt.Println(bag01.Resolve(w, b))
	}

	// {
	// 	data := []int{1,2,3,4}
	// 	arrange.Arrange(data, 0)
	// }

	{
		g := graph.NewGraph(5, graph.DIRECTION_BOTH)
		g.Insert(0, 1, 10)
		g.Insert(0, 3, 4)
		g.Insert(0, 4, 12)
		g.Insert(1, 3, 8)
		g.Insert(1, 4, 5)
		g.Insert(1, 2, 15)
		g.Insert(2, 4, 30)
		g.Insert(2, 3, 7)
		g.Insert(3, 4, 6)

		s := 0
		city := []int{s}
		for i := 0; i < g.V; i++ {
			if s != i {
				city = append(city, i)
			}
		}

		curcost := 0
		best := graph.MAXDIST
		bestpath := make([]int, len(city))
		tsp.TSP(g, s, city, 1, curcost, &best, &bestpath)
		fmt.Println(best)
		fmt.Println(bestpath)
	}

	{
		g := graph.NewGraph(5, graph.DIRECTION_BOTH)
		g.Insert(0, 1, 10)
		g.Insert(0, 3, 4)
		g.Insert(0, 4, 12)
		g.Insert(1, 3, 8)
		g.Insert(1, 4, 5)
		g.Insert(1, 2, 15)
		g.Insert(2, 4, 30)
		g.Insert(2, 3, 7)
		g.Insert(3, 4, 6)

		color := make([]int, g.V)
		colorgraph.Color(g, 4, 0, color)
	}

	{
		editdist.EditDist(0, 0, 0)
		fmt.Println(editdist.MinDist)
	}
}

func dpTest() {
	{
		dpbag01.InitBag()
		fmt.Println(dpbag01.Bag01_2())
	}
	{
		fmt.Println(yanghui.DpYanghui())
	}
	{
		fmt.Println(coin.DpCoin())
	}
	{
		fmt.Println(arraymaxsum.DpArrayMaxSum2())
	}
	{
		fmt.Println(dpeditdist.DpEditDist())
	}
	{
		fmt.Println(lcs.LCS())
	}
	{
		fmt.Println(lis.LIS())
	}
}

func main() {
	// sortTest()
	// binSearchTest()
	// hashMapTest()
	// arrangeTest()
	// cellSplitTest()
	// linkedlistTest()
	// skipListTest()
	// binSarchTreeTest()
	treeTest()
	// graphTest()
	// stringMatchTest()
	// trieTest()
	// greedyTest()
	// dcTest()
	// btTest()
	// dpTest()
}
