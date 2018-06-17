package algo

import "fmt"

type Mammal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "I'm a dog"
}

type HotDog struct {
	Hot *Hot
	Dog *Dog
}

type Hot struct{}

func (h Hot) Speak() string {
	return "I'm hot"
}

func (h *HotDog) Speak() string {
	return "I'm a hotdog"
}

func MammalStuff() {
	var d Dog
	fmt.Println(d.Speak())

	var m Mammal
	m = Dog{}
	fmt.Println(m.Speak())

	var h HotDog
	fmt.Println(h.Speak())
}
