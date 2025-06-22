package structure

type ListNode struct {
	Val  int
	Next *ListNode
}

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
