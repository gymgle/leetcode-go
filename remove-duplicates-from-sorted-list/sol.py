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
        num = p.val
        while p.next:
            if p.next.val == num:
                p.next = p.next.next
            else:
                num = p.next.val
                p = p.next

        return head
