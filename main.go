package main

import "fmt"

//selection sort
//takes the smallest element in a given array and adds it to a new array
// should be n^2 time
// func selectionSort(arr []int) []int {
// 	var sorted []int
// 	for j = 0; j < len(arr) - 1; j++ {
// 		iMin = j
// 		for i = j+1; i < len(arr); i++ {
// 			if a[i] < a[j] {
// 				a[i] = a[j]
// 			}
// 		}
// 	}

// }

//merge sort
//split the array in two halves, sort each array, merge them

// func sort(arr []int) []int {
// 	for i := 0; i < len(arr); i++ {
// 		for j := 0; j < len(arr); j++ {
// 			if arr[i] < arr[j] {
// 				arr[i], arr[j] = arr[j], arr[i]
// 			}
// 		}
// 	}

// 	return arr
// }

func merge(left, right []int) []int {
	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	fmt.Println(slice)
	return slice
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]

	fmt.Println(merge(mergeSort(left), mergeSort(right)))
	return merge(mergeSort(left), mergeSort(right))
}

func main() {
	ar := []int{5, 3, 4, 1, 2}
	mergeSort(ar)
}
