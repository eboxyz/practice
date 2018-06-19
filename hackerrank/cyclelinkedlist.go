package hackerrank

type LinkedNode struct {
	Data int
	Next *LinkedNode
}

func checkCycle(slow *LinkedNode, fast *LinkedNode) bool {
	if fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	} else {
		checkCycle(slow, fast)
	}

	return false
}

func HasCycle(head *LinkedNode) bool {
	if head.Next == nil {
		return true
	}

	slow := head
	fast := head

	return checkCycle(slow, fast)

}
