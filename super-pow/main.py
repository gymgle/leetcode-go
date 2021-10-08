# coding: utf-8


class Solution(object):
    def superPow(self, a, b):
        """
        :type a: int
        :type b: List[int]
        :rtype: int
        """
        if not b:
            return 1
        last = b.pop()
        p1 = _pow(a, last)
        p2 = _pow(self.superPow(a, b), 10)
        return (p1 * p2) % 1337


def _pow(a, b):
    if not b:
        return 1
    rst = 1
    a = a % 1337
    for _ in range(b):
        rst = (rst * a) % 1337
    return rst


if __name__ == '__main__':
    s = Solution()
    print s.superPow(2, [1, 0])

