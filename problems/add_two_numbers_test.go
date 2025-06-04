package problems

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func NewLinkedList(vals ...int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := &ListNode{Val: vals[0]}
	current := head

	for _, val := range vals[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}

	return head
}

func (l *ListNode) EqualValues(o *ListNode) bool {
	n1, n2 := l, o
	for {
		// Either both are nil or they're the same node
		if n1 == n2 {
			return true
		}

		// Can both be nil as we would've returned above
		if n1 == nil || n2 == nil {
			return false
		}

		if n1.Val != n2.Val {
			return false
		}

		n1, n2 = n1.Next, n2.Next
	}
}

func TestAddTwoNumbers(t *testing.T) {
	cases := []struct {
		l1       *ListNode
		l2       *ListNode
		expected *ListNode
	}{
		{
			NewLinkedList(2, 4, 3),
			NewLinkedList(5, 6, 4),
			NewLinkedList(7, 0, 8),
		},
		{
			NewLinkedList(0),
			NewLinkedList(0),
			NewLinkedList(0),
		},
		{
			NewLinkedList(9, 9, 9, 9, 9, 9, 9),
			NewLinkedList(9, 9, 9, 9),
			NewLinkedList(8, 9, 9, 9, 0, 0, 0, 1),
		},
	}

	for _, testCase := range cases {
		t.Run(fmt.Sprintf("%+v", testCase), func(tt *testing.T) {
			assert.True(t, testCase.expected.EqualValues(AddTwoNumbers(testCase.l1, testCase.l2)))
		})
	}
}

func TestNewLinkedList(t *testing.T) {
	l := NewLinkedList(0, 1, 2)
	assert.Equal(t, 0, l.Val)
	l = l.Next
	assert.Equal(t, 1, l.Val)
	l = l.Next
	assert.Equal(t, 2, l.Val)
	l = l.Next
	assert.Nil(t, l)
}

func TestNewLinkedList_Empty(t *testing.T) {
	l := NewLinkedList()
	assert.Nil(t, l)
}

func TestListNode_EqualValues(t *testing.T) {
	assert.True(t, NewLinkedList(1, 2, 3).EqualValues(NewLinkedList(1, 2, 3)))
	assert.False(t, NewLinkedList(1, 2, 3).EqualValues(NewLinkedList(4, 5)))
}
