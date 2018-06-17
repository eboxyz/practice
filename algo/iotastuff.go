package algo

import (
	"fmt"
)

type ByteSize float64

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota) // 1 << (10*1)
	MB                             // 1 << (10*2)
	GB                             // 1 << (10*4)
	TB
	PB
	EB
	ZB
	YB
)

const (
	Apple, Banana = iota + 1, iota + 2
	Cherimoya, Durian
	Elderberry, Fig
)

func IotaPractice() {
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)

	fmt.Println(Apple)
	fmt.Println(Banana)

	fmt.Println(Cherimoya, Durian)
}
