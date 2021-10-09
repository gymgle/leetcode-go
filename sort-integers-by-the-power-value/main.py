# coding: utf-8


class Solution(object):
    def getKth(self, lo, hi, k):
        """
        :type lo: int
        :type hi: int
        :type k: int
        :rtype: int
        """
        v = list(range(lo, hi + 1))
        v.sort(key=lambda x: (_power(x), x))
        return v[k - 1]


def _power(n):
    if n == 1:
        return 0
    if n % 2 == 0:
        return _power(n / 2) + 1
    else:
        return _power(n * 3 + 1) + 1


if __name__ == '__main__':
    s = Solution()
    print s.getKth(12, 15, 2)

