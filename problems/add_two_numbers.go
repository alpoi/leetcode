package problems

// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order, and each of their nodes contains
// a single digit. Add the two numbers and return the sum as a linked list.
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	n1, n2 := l1, l2
	carry := 0
	head := &ListNode{Val: 0}
	current := head

	for n1 != nil || n2 != nil || carry != 0 {
		val := carry

		if n1 != nil {
			val += n1.Val
			n1 = n1.Next
		}

		if n2 != nil {
			val += n2.Val
			n2 = n2.Next
		}

		carry = val / 10
		current.Next = &ListNode{Val: val % 10}
		current = current.Next
	}

	return head.Next
}
