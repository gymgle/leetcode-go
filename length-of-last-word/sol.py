class Solution:
    def lengthOfLastWord(self, s: str) -> int:
        count = 0
        length = len(s)
        start = length - 1
        while start >= 0:
            if s[start] == ' ':
                start -= 1
                continue
            break

        while start >= 0:
            if s[start] != ' ':
                count += 1
                start -= 1
                continue
            break
        return count
