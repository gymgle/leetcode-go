class Solution:
    def mySqrt(self, x: int) -> int:
        if x == 1:
            return 1

        i, j = 0, x
        while j - i > 1:
            m = (j - i) // 2 + i
            if x//m < m:
                j = m
            else:
                i = m
        return i
