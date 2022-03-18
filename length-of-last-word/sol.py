class Solution:
    def lengthOfLastWord(self, s: str) -> int:
        cnt, p = 0, len(s) - 1
        while s[p] == ' ':
            p -= 1
        while p >= 0 and s[p] != ' ':
            p -= 1
            cnt += 1
        return cnt
