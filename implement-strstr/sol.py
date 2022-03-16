class Solution:
    def strStr(self, haystack: str, needle: str) -> int:
        if not needle:
            return 0
        if not haystack:
            return -1

        s1 = len(haystack)
        s2 = len(needle)
        if s1 < s2:
            return -1

        return haystack.find(needle)
