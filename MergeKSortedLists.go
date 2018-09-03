package leetcode

/*
Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.

Example:

Input:
[
  1->4->5,
  1->3->4,
  2->6
]
Output: 1->1->2->3->4->4->5->6
*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	i := 0
	vm := make(map[int]*ListNode)
	var min *ListNode
	for {
		minIndex := 0
		for j := 0; j < len(lists); j++ {
			if lists[j] != nil {
				min = lists[j]
				minIndex = j
			}
		}
		if min == nil {
			break
		}
		for j := 0; j < len(lists); j++ {
			if lists[j] != nil && min.Val > lists[j].Val {
				min = lists[j]
				minIndex = j
			}
		}
		vm[i] = min
		i++
		min = nil
		lists[minIndex] = lists[minIndex].Next
	}
	var result, p *ListNode
	if node, ok := vm[0]; ok {
		result = node
		p = result
		p.Next = nil
		for j := 1; j < i; j++ {
			p.Next = vm[j]
			p = p.Next
			p.Next = nil
		}
	}
	return result
}

func MergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	var min *ListNode
	var result, p *ListNode
	for {
		minIndex := 0
		for j := 0; j < len(lists); j++ {
			if lists[j] != nil {
				min = lists[j]
				minIndex = j
				break 
			}
		}
		if min == nil {
			break
		}
		for j := 0; j < len(lists); j++ {
			if lists[j] != nil && min.Val > lists[j].Val {
				min = lists[j]
				minIndex = j
			}
		}
		if result == nil {
			result = min
			p = result
			p.Next = nil
		} else {
			p.Next = min
			p = p.Next
			p.Next = nil
		}
		min = nil
		lists[minIndex] = lists[minIndex].Next
	}
	return result
}

func listEqual(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}
