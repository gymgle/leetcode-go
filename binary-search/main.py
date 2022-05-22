from typing import List


class Solution:
    def search(self, nums: List[int], target: int) -> int:
        if len(nums) == 0:
            return -1

        i, j = 0, len(nums) - 1
        while i <= j:
            mid = i + (j - i) // 2
            if target > nums[mid]:
                i = mid + 1
            elif target < nums[mid]:
                j = mid - 1
            else:
                return mid

        return -1
