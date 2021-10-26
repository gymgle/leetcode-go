# coding: utf-8

class Solution(object):
    def isPalindrome(self, x):
        """
        :type x: int
        :rtype: bool
        """
        """ solution 1. turn number to string, need a string space
        if x < 0:
            return False
        s = str(x)
        for i in range(len(s) // 2):
            if s[i] != s[-1 - i]:
                return False
        return True
        """
        # solution 2
        


if __name__ == '__main__':
    sol = Solution()
    print(sol.isPalindrome(12321))
    print(sol.isPalindrome(123))
