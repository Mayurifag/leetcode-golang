package leetcode

type ListNode struct {
	Val int
	Next *ListNode
}

// Space - O(1) - nothing to save
// Time - O(n) - iterate through each node
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
			return head
	}
	if head.Val == val {
			return removeElements(head.Next, val)
	}
	head.Next = removeElements(head.Next, val)
	return head
}
