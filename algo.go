package main

import (
	"math/rand"
	"algo/bin_search"
	"algo/cmpsort"
	"algo/hashmap"
	"algo/linkedlist"
	"algo/recursive/arrange"
	"fmt"
	"sort"
)

func sortTest() {
	x := make([]int, 0)
	n := 100
	for i:=0; i<n; i++ {
		x = append(x, rand.Intn(n))
	}
	x = cmpsort.MergeSort(x)
	// cmpsort.HeapSort(x)
	// cmprsort.QuicSort(x)
	for _,v := range x {
		fmt.Println(v)
	}
}

func binSearchTest() {
	x := make([]int, 0)
	n := 1000
	for i:=0; i<n; i++ {
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
	arrange.Arrange(A, 0, len(A))
}

func linkedlistTest() {
	head := linkedlist.InitLinkedList()
	linkedlist.Range(head)
	head = linkedlist.Reverse(head)
	linkedlist.Range(head)
}

func main () {
	sortTest()
	binSearchTest()
	hashMapTest()
	arrangeTest()
	linkedlistTest()
}
