import collections


class Solution:
    def longestPalindrome(self, s: str) -> int:
        ans = 0
        single = False
        counter = collections.Counter(s)
        for v in counter.values():
            ans += v // 2 * 2
            if not single and v % 2 == 1:
                single = True

        return ans + 1 if single else ans
