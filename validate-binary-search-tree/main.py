from asyncio.windows_events import NULL
from logging import NullHandler
from turtle import isvisible
from typing import Optional


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def isValidBST(self, root: Optional[TreeNode]) -> bool:
        self.pre = None

        def _dfs(root: Optional[TreeNode]) -> bool:
            if not root:
                return True

            # left tree
            if not _dfs(root.left):
                return False

            # in order process
            if self.pre is not None and root.val <= self.pre:
                return False
            self.pre = root.val

            # right tree
            return _dfs(root.right)

        return _dfs(root)
