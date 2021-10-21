# coding: utf-8

"""
使用字典查询数字是否存在，查询时间复杂度可降为 O(1)

如何把整体时间复杂度降为 O(n):
如果已知有一个连续序列 x x+1 x+2 ... x+n 长度为 n+1，对于 x ~ x+n 之间的数字，已经没有遍历的必要，长度肯定小于 n+1。
因此，在遍历数组 nums 时，如果发现当前数字的前驱数字存在，则可以直接跳过。

时间复杂度分析：虽然最终是双层循环，外层循环时间复杂度为 O(n)，但内层循环并非完全遍历，只有当下一个数是连续序列的第一个数时才会进入循环，
也就意味着每一个数只会被遍历一次，因此整体时间复杂度为 O(n)
"""


class Solution(object):
    def longestConsecutive(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        m = dict()
        for i in nums:
            m[i] = True

        max_length = 0
        for i in nums:
            if m.get(i - 1, False):
                continue
            cur_length = 1
            for j in range(i + 1, len(nums) + 1):
                if not m.get(j, False):
                    break
                cur_length += 1
            max_length = cur_length if cur_length > max_length else max_length
        return max_length


if __name__ == '__main__':
    sol = Solution()
    print(sol.longestConsecutive([1, 3, 2, 4]))  # 4
    print(sol.longestConsecutive([1, 3, 2, 100, 200, 5, 4]))  # 5
    print(sol.longestConsecutive([0, 3, 7, 2, 5, 8, 4, 6, 0, 1]))  # 9

