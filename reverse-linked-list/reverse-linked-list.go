package leetcode

/**
 * 206. Reverse Linked List
 *
 * Reverse a singly linked list.
 *
 * Example:
 * Input: 1->2->3->4->5->NULL
 * Output: 5->4->3->2->1->NULL
 *
 * https://leetcode.com/problems/reverse-linked-list/
 */

// 主要考察翻转的写法，用 go 语言可以精简到一行，不需要 c/c++ 使用中间变量

// ListNode definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 迭代解法
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

/**
 * 递归解法
 * 对于每一个链表节点，我们考虑先将其后续的链表进行反转。
 *
 * 1->2->3->4->5->NULL 假设我们已经将后面两个节点反转了得到了
 * 1—>2->3->4<-5
 *      cur    p
 * 那么我们执行cur->next->next = cur,cur->next = NULL可以得到
 * 1—>2->3<-4<-5
 * 递归进行即可
 */

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead
}
