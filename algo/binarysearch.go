package algo

//binary search, searches through an array for a specific element

//Search ... ok?
func Search(arr []int, num int) bool {
	start := 0
	end := len(arr) - 1
	middle := (start + end) / 2
	var ok bool

	for i := start; i < end; i++ {
		if arr[middle] == num {
			ok = true
		} else if arr[middle] < num {
			start = middle + 1
		} else if arr[middle] > num {
			end = middle - 1
		}
	}

	return ok
}
