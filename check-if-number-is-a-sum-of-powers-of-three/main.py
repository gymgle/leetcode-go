# coding: utf-8


class Solution(object):
    def checkPowersOfThree(self, n):
        """
        :type n: int
        :rtype: bool
        """
        while n > 0:
            if n % 3 == 2:
                return False
            print(n)
            # n /= 3  # python 2
            n //= 3  # python 3
        return True


if __name__ == '__main__':
    sol = Solution()
    print(sol.checkPowersOfThree(21))
    print(sol.checkPowersOfThree(91))
