# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution(object):
    def inorderTraversal(self, root):
        """
        :type root: TreeNode
        :rtype: List[int]
        """
        rst = list()

        def dfs(root):
            if not root:
                return
            dfs(root.left)
            rst.append(root.val)
            dfs(root.right)

        dfs(root)
        return rst


if __name__ == '__main__':

    root = TreeNode(1)
    root.right = TreeNode(2)
    root.right.left = TreeNode(3)

    s = Solution()
    print(s.inorderTraversal(root))
