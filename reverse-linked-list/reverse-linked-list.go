package leetcode

/**
 * 206. Reverse Linked List
 *
 * Reverse a singly linked list.
 *
 * Example:
 * Input: 1->2->3->4->5->NULL
 * Output: 5->4->3->2->1->NULL
 */

// 主要考察翻转的写法，用 go 语言可以精简到一行，不需要 c/c++ 使用中间变量

// ListNode definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		cur.Next, pre, cur = pre, cur, cur.Next
		/* 等效以下代码
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
		*/
	}
	return pre
}
