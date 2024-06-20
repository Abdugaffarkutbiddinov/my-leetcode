package convert_binary_number_in_a_linked_list_to_integer

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) (decimal int) {

	for head != nil {
		decimal = (decimal * 2) + head.Val
		head = head.Next
	}

	return
}
