class Solution:
    def searchInsert(self, nums: list[int], target: int) -> int:
        left = 0
        right = len(nums) - 1
        while left < right:
            middle = left + (right - left) // 2
            if nums[middle] > target:
                right = middle - 1
            elif nums[middle] < target:
                left = middle + 1
            else:
                return middle

        return left if nums[left] >= target else left + 1


if __name__ == '__main__':
    s = Solution()
    print(s.searchInsert([1, 2, 3, 4, 5], 7))
