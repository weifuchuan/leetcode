package leetcode

/*
Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.

k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.

Example:

Given this linked list: 1->2->3->4->5

For k = 2, you should return: 2->1->4->3->5

For k = 3, you should return: 3->2->1->4->5

Note:

Only constant extra memory is allowed.
You may not alter the values in the list's nodes, only nodes itself may be changed.
*/
func ReverseKGroup(head *ListNode, k int) *ListNode {
	var begin, end, lastEnd *ListNode
	begin = head
	p := head
	isFirst := true
	for i := 0; p != nil; {
		if i == k-1 {
			end = p
			p = p.Next
			begin, end = reverse(begin, end)
			if isFirst {
				isFirst = false
				head = begin
			}
			if lastEnd != nil {
				lastEnd.Next = begin
			}
			begin = end.Next
			lastEnd = end
			i = 0
		} else {
			p = p.Next
			i++
		}
	}
	return head
}

func reverse(begin *ListNode, end *ListNode) (begin2 *ListNode, end2 *ListNode) {
	if begin == end {
		return begin, end
	}
	tail := end.Next
	p := begin
	for {
		if p == end {
			break
		}
		if end.Next == tail {
			end.Next = p
			p = p.Next
		} else {
			q := p
			p = p.Next
			q.Next = end.Next
			end.Next = q
		}
	}
	begin.Next = tail
	return end, begin
}
