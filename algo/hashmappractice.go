package algo

import "fmt"

//HashMap ... ok
func HashMap() {
	// var newMap map[string]int <- alternate
	newMap := make(map[string]int)
	newMap["hello"] = 1
	newMap["goodbye"] = 2

	intMap := make(map[int]string)
	intMap[-1] = "blah"
	intMap[0] = "hullo"
	intMap[1] = "hahahahaahha"
	intMap[2] = "hoohaha"

	stringMap := map[string]string{
		"hello":   "world",
		"goodbye": "everyone",
	}
	fmt.Println(stringMap)

	for i := -1; i < len(intMap); i = i + 2 {
		fmt.Println(intMap[i])
	}

	fmt.Println(len(newMap))

	fmt.Println(newMap)
}
