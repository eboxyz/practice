package linkedlist

// type Node struct {
// 	next, prev *Node
// 	value      interface{}
// 	list       *List
// }

// func (n *Node) Next() *Node {
// 	if p := n.next; n.list != nil && p != &n.list.root {
// 		return p
// 	}
// 	return nil
// }

// func (n *Node) Prev() *Node {
// 	if p := n.prev; n.list != nil && p != &n.list.root {
// 		return p
// 	}
// 	return nil
// }

// type List struct {
// 	root   Node //this is a sentinel node, so we use it as a marker of sorts. It only has a pointer, pointer to prev and pointer to next
// 	length int  //current list length, does not include the sentinel node
// }

// func (l *List) Init() *List {
// 	l.root.next = &l.root
// 	l.root.prev = &l.root
// 	l.length = 0
// 	return l
// }

// func New() *List { return new(List).Init() }

// func (l *List) Len() int { return l.length }

// func (l *List) Front() *Node {
// 	if l.length == 0 {
// 		return nil
// 	}
// 	return l.root.next
// }

// //this list isn't doubly linked, so we need to
// //look through the whole list and find the last element
// // this should be O(N) where N is the number of nodes in the list
// func (l *List) Back() *Node {
// 	if l.length == 0 {
// 		return nil
// 	}
// 	return l.root.prev
// }

// func (l *List) lazyInit() {
// 	if l.root.next == nil {
// 		l.Init()
// 	}
// }

// func (l *List) insert(n, at *Node) *Node {
// 	// instantialize position of new node
// 	new := at.next
// 	// change pointer of inserted at to be the new node
// 	at.next = n
// 	// update inserted at pointer to point at where it was inserted
// 	n.prev = at
// 	// update inserted at pointer to point at its next oject
// 	n.next = new
// 	// set the previous to point at the newly inserted node
// 	new.prev = n
// 	// add the list pointer to the node
// 	n.list = l
// 	// up list length
// 	l.length++
// 	return n
// }

// //convenience wrapper for insert()
// func (l *List) insertValue(v interface{}, at *Node) *Node {
// 	return l.insert(&Node{value: v}, at)
// }

// func (l *List) remove(n *Node) *Node {
// 	// set the previous node's 'next' to be the next node
// 	n.prev.next = n.next
// 	// set next node's 'prev' to be the previous node
// 	n.next.prev = n.prev
// 	// remove all values from node 'n'
// 	n.next = nil
// 	n.prev = nil
// 	n.list = nil
// 	// down list length
// 	l.length--
// 	return n
// }

// func (l *List) Remove(n *Node) interface{} {
// 	if n.list == l {
// 		l.remove(n)
// 	}
// 	return n.value
// }

// func (l *List) PushFront(v interface{}) *Node {
// 	l.lazyInit()
// 	return l.insertValue(v, &l.root)
// }

// func (l *List) PushBack(v interface{}) *Node {
// 	l.lazyInit()
// 	return l.insertValue(v, l.root.prev)
// }

// func (l *List) InsertBefore(v interface{}, mark *Node) *Node {
// 	if mark.list != l {
// 		return nil
// 	}
// 	return l.insertValue(v, mark.prev)
// }

// func (l *List) InsertAfter(v interface{}, mark *Node) *Node {
// 	if mark.list != l {
// 		return nil
// 	}
// 	return l.insertValue(v, mark)
// }

// func (l *List) MoveToFront(n *Node) {
// 	if n.list != l || l.root.next == n {
// 		return
// 	}
// 	l.insert(l.remove(n), &l.root)
// }

// func (l *List) MoveToBack(n *Node) {
// 	if n.list != l || l.root.prev == n {
// 		return
// 	}
// 	l.insert(l.remove(n), l.root.prev)
// }

// func (l *List) MoveBefore(n, mark *Node) {
// 	if n.list != l || n == mark || mark.list != l {
// 		return
// 	}
// 	l.insert(l.remove(n), mark.prev)
// }

// func (l *List) MoveAfter(n, mark *Node) {
// 	if n.list != l || n == mark || mark.list != l {
// 		return
// 	}
// 	l.insert(l.remove(n), mark)
// }
