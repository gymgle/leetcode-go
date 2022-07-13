from typing import List


# Definition for a Node.
class Node:
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children


class Solution:
    def preorder(self, root: 'Node') -> List[int]:
        rst = list()

        def dfs(root: 'Node'):
            if not root:
                return
            rst.append(root.val)
            for c in root.children:
                dfs(c)

        dfs(root)
        return rst
