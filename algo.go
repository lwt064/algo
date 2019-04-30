package main

import (
	"algo/bin_search"
	"algo/bin_search_tree"
	"algo/cmpsort"
	hashMap "algo/hashmap"
	"algo/linkedlist"
	"algo/recursive/arrange"
	"algo/recursive/cellsplit"
	"algo/skiplist"
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
	// // cmprsort.QuicSort(x)

	fmt.Println("kth: ", cmpsort.SelectK(x, 18))

	// cmpsort.InsertSort(x)
	// for _,v := range x {
	// 	fmt.Println(v)
	// }
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

func arrangeTest() {
	A := []int{1, 2, 3}
	arrange.Arrange(A, len(A), 0)
}

func linkedlistTest() {
	head := linkedlist.InitLinkedList()
	linkedlist.Range(head)
	head = linkedlist.Reverse(head)
	linkedlist.Range(head)
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

func main() {
	// sortTest()
	// binSearchTest()
	// hashMapTest()
	// arrangeTest()
	cellSplitTest()
	// linkedlistTest()
	// skipListTest()
	// binSarchTreeTest()
}
