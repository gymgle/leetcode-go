/**
 * You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
 * You may assume the two numbers do not contain any leading zero, except the number 0 itself.
 * Example
 * Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
 * Output: 7 -> 0 -> 8
 * Explanation: 342 + 465 = 807.
 *
 * 思路时按照竖式相加去计算，保留进位。
 * 需要注意的是，两个链表长度可能不同，所以不能直接相加。当最后一对 Nodes 相加有进位时，需要把进位 1 保存到新的一个 Node 中作为尾结点。
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 进位值, 只可能为0或1
	promotion := 0

	// 结果表的头结点
	var head *ListNode
	// 保存结果表的尾结点
	var rear *ListNode
	for nil != l1 || nil != l2 {
		sum := 0
		if nil != l1 {
			sum += l1.Val
			l1 = l1.Next
		}
		if nil != l2 {
			sum += l2.Val
			l2 = l2.Next
		}

		sum += promotion
		promotion = 0

		if sum >= 10 {
			promotion = 1
			sum = sum % 10
		}

		node := &ListNode{
			sum,
			nil,
		}

		if nil == head {
			head = node
			rear = node
		} else {
			rear.Next = node
			rear = node
		}

	}

	if promotion > 0 {
		rear.Next = &ListNode{
			promotion,
			nil,
		}
	}

	return head
}