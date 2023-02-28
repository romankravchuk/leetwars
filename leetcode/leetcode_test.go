package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getNodeValues(node *ListNode) []int {
	out := make([]int, 0)
	curr := node
	for curr != nil {
		out = append(out, curr.Val)
		curr = curr.Next
	}
	return out
}

func Test_AddTwoNumbers(t *testing.T) {
	t.Run("case-1", func(t *testing.T) {
		l1 := &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		}
		l2 := &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		}
		l3 := AddTwoNumbers(l1, l2)
		assert.Equal(t, []int{7, 0, 8}, getNodeValues(l3))
	})
}
