package helpers

func Swap(arr []int, i, j int) {
	k := arr[i]
	arr[i] = arr[j]
	arr[j] = k
}
