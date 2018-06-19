package linkedlist

import "fmt"

type Node struct {
	next, prev *Node
	value      interface{}
	List       *List
}

type List struct {
	Head   *Node
	Length int
}

func New() *List {
	l := new(List)
	l.Head = &Node{}
	l.Length = 0
	return l
}

//Insert ... ok?
func (l *List) Insert(n, at *Node) *List {
	fmt.Println(l)
	//we're inserted it at that position, so n will take at's previous
	n.prev = at.prev
	//moving at back one, so n's next will be at
	n.next = at
	//at's previous needs to be n
	at.next = n
	n.List = l
	l.Length++
	return l
}
