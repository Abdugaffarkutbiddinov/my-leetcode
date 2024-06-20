package reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

// So when we allocate head.next to prev we make head look to another prevList and take he head of it
func reverseList(head *ListNode) *ListNode {
	var next, prev *ListNode

	for head != nil {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}
