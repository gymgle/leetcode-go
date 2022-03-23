from typing import List


class Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        """
        Do not return anything, modify nums1 in-place instead.
        """
        p = m + n - 1
        m, n = m - 1, n - 1
        while m >= 0 and n >= 0:
            if nums1[m] > nums2[n]:
                nums1[p] = nums1[m]
                m -= 1
            else:
                nums1[p] = nums2[n]
                n -= 1
            p -= 1

        while n >= 0:
            nums1[p] = nums2[n]
            p, n = p - 1, n - 1


if __name__ == '__main__':
    s = Solution()
    n1 = [1, 3, 5, 0, 0, 0]
    s.merge(n1, 3, [2, 4, 6], 3)
    print(n1)
