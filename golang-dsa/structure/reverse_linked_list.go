package structure

func reverseList(head *ListNode) *ListNode {
	var newList *ListNode  // nil
	var nextNode *ListNode // nil

	for head != nil {
		nextNode = head.Next
		head.Next = newList
		newList = head
		head = nextNode
	}

	return newList
}
