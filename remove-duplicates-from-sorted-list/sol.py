from typing import Optional


# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def deleteDuplicates(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if not head:
            return head

        p = head
        while p.next:
            if p.next.val == p.val:
                p.next = p.next.next
            else:
                p = p.next

        return head


if __name__ == '__main__':
    head = ListNode(1, None)
    head.next = ListNode(1, None)
    head.next.next = ListNode(2, None)
    head.next.next.next = ListNode(2, None)

    s = Solution()
    s.deleteDuplicates(head)
    while head:
        print(head.val)
        head = head.next
