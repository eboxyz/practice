package main

import (
	linkedlist "algo-boi/algo/linked-lists"
	"fmt"
)

/* for graph tree node search stuff */
// var nodes = map[int][]int{
// 	1:  []int{2, 3, 4},
// 	2:  []int{1, 5, 6},
// 	3:  []int{1},
// 	4:  []int{1, 7, 8},
// 	5:  []int{2, 9, 10},
// 	6:  []int{2},
// 	7:  []int{4, 11, 12},
// 	8:  []int{4},
// 	9:  []int{5},
// 	10: []int{5},
// 	11: []int{7},
// 	12: []int{7},
// }

func main() {
	ll := linkedlist.New()
	fmt.Println(ll)
	n := &linkedlist.Node{}
	fmt.Println(n)
	l1 := ll.Insert(n, ll.Head)
	fmt.Println(l1)
	// ar := []int{5, 3, 4, 1, 2}
	// mergesort.MergeSort(ar)
	// var one float64
	// one = 3141592653589793238462643383279502884197169399375105820974944592

	// fmt.Println(ll)
	// coursera.Solve(1, 2)

	// algo.IotaPractice()
	// algo.MammalStuff()
	// visited := []int{}
	// search.BFS(1, nodes, func(node int) {
	// 	visited = append(visited, node)
	// })
	// fmt.Println(visited)
	// ar := []int{5, 4, 3, 2, 1}
	// fmt.Println(algo.SelectionSort(ar))
	// var acc int
	// algo.Divide(2, 3, acc)

}
