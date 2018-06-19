package search

func BFS(start int, nodes map[int][]int, fn func(int)) {
	//
	frontier := []int{start}
	visited := map[int]bool{}
	next := []int{}

	for 0 < len(frontier) {
		next = []int{}
		for _, node := range frontier {
			visited[node] = true
			fn(node)
			for _, n := range bfsFrontier(node, nodes, visited) {
				next = append(next, n)
			}
		}
		frontier = next
	}
}

func bfsFrontier(node int, nodes map[int][]int, visited map[int]bool) []int {
	next := []int{}
	iter := func(n int) bool { _, ok := visited[n]; return !ok }
	// for _, n := range nodes[node] {
	for i := 0; i < len(nodes); i++ {
		if iter(i) {
			next = append(next, i)
		}
	}
	// if iter(n) {
	// 	next = append(next, n)
	// }
	return next
}
