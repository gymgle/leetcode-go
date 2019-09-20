package leetcode

/*
 * 24. Swap Nodes in Pairs
 * Given a linked list, swap every two adjacent nodes and return its head.
 * You may not modify the values in the list's nodes, only nodes itself may be changed.
 * Example:
 * Given 1->2->3->4, you should return the list as 2->1->4->3.
 *
 * https://leetcode.com/problems/swap-nodes-in-pairs/
 */

// ListNode definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 使用递归实现
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := head.Next
	newHead.Next, head.Next = head, swapPairs(newHead.Next)

	return head
}

/**
 * 使用迭代实现
 * 对于 1->2->3->4->5 的链表，
 * 使用 a b 分别指向待交换的两个结点，pre 指向 a 的前结点，
 *  0──>1->2->3->4->5
 * pre  a  b
 *
 *      ┌─────┐
 *  0   1<-2  3->4->5
 *  └──────┘
 *     pre    a  b
 */
func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := head.Next
	pre := new(ListNode)
	pre.Next = head

	for pre.Next != nil && pre.Next.Next != nil {
		a := pre.Next
		b := a.Next

		pre.Next, a.Next, b.Next = b, b.Next, a

		pre = a
	}

	return newHead
}
