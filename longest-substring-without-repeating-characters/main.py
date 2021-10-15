# coding: utf-8

"""
滑动窗口实现，窗口的实现：用一个 set 保存未重复出现的字符
"""


class Solution(object):
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        if not s:
            return 0
        n = len(s)
        left = 0  # left side of window
        max_len = 0
        cur_len = 0
        lookup = set()
        for i in range(n):
            cur_len += 1
            while s[i] in lookup:
                lookup.remove(s[left])
                left += 1
                cur_len -= 1
            lookup.add(s[i])
            max_len = cur_len if cur_len > max_len else max_len
        return max_len


if __name__ == '__main__':
    sol = Solution()
    print sol.lengthOfLongestSubstring('pwwkew')

