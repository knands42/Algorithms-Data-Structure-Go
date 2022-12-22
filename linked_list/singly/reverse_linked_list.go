package singly

func ReverseLinkedList[T comparable](head *Node[T]) *Node[T] {
	var prev *Node[T]
	var current *Node[T] = head
	var next *Node[T]

	for current != nil {
		next = current.next
		current.next = prev

		prev = current
		current = next
	}

	return prev
}
