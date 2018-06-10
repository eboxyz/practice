package algo

//binary search, searches through an array for a specific element

//Search ... ok?
func Search(arr []int, num int) bool {
	middle := len(arr) / 2
	left := arr[:middle]
	right := arr[middle:]

	var ok bool

	if num < arr[middle] {
		for i := 0; i < len(left); i++ {
			if left[i] == num {
				ok = true
			}
		}
	} else if num >= arr[middle] {
		for i := 0; i < len(right); i++ {
			if right[i] == num {
				ok = true
			}
		}
	} else {
		ok = false
	}

	return ok
}
