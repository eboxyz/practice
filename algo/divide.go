package algo

import "fmt"

//Divide ... ok?
func Divide(a, b, acc int) int {
	if a == 0 {
		fmt.Println("cannot divide by zero")
	}
	if b == 0 {
		return 0
	}
	//we want to count the number of times that
	//we can subtract a from b
	b = b - a
	acc++
	Divide(a, b, acc)

	return acc
}
