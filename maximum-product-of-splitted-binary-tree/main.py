# leetcode 1339

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def maxProduct(self, root):
        sum_list = list()

        def get_sum(root):
            if not root:
                return 0
            left_sum = get_sum(root.left)
            right_sum = get_sum(root.right)
            sum_list.append(left_sum)
            sum_list.append(right_sum)
            return root.val + right_sum + left_sum

        total = get_sum(root)
        mod = 10**9 + 7
        half_total = total / 2
        print(half_total)
        print(total)
        print(sum_list)
        prop = half_total
        min_sub = half_total  # 挑选一个离中间值最近的一个数字
        for i in sum_list:
            if abs(half_total - i) < min_sub:
                min_sub = abs(half_total - i)
                prop = i
        return (prop * (total - prop)) % mod


if __name__ == '__main__':
    t = TreeNode(2)
    t.right = TreeNode(4)
    t.left = TreeNode(5)

    s = Solution()
    print(s.maxProduct(t))
