from typing import Optional


# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        prev, curr = None, head
        while curr:
            # tmp = curr.next
            # curr.next = prev
            # prev = curr
            # curr = tmp
            curr.next, prev, curr = prev, curr, curr.next
        return prev


if __name__ == '__main__':
    n3 = ListNode(3, None)
    n2 = ListNode(2, n3)
    head = ListNode(1, n2)
    sol = Solution()
    sol.reverseList(head)
