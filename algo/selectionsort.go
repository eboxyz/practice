package algo

//SelectionSort ... ok
func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		min := 1
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}
