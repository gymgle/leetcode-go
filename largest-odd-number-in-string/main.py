class Solution(object):
    def largestOddNumber(self, num):
        """
        :type num: str
        :rtype: str
        """
        idx = len(num) - 1
        while idx >= 0:
            if int(num[idx]) % 2 == 1:
                return num[:idx + 1]
            idx -= 1
        if idx < 0:
            return ''


if __name__ == '__main__':
    s = Solution()
    print(s.largestOddNumber('2'))
