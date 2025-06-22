package structure

func middleNode(head *ListNode) *ListNode {
	aHead := head

	for aHead != nil && aHead.Next != nil {
		aHead = aHead.Next.Next
		head = head.Next
	}

	return head
}
