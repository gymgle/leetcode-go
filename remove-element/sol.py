class Solution:
    def removeElement(self, nums: list[int], val: int) -> int:
        if not nums:
            return 0

        last = len(nums) - 1
        while last > 0:
            if nums[last] == val:
                last -= 1
            else:
                break
        if last == 0 and nums[last] == val:
            return 0
        i = last - 1
        while i >= 0:
            if nums[i] == val:
                nums[i] = nums[last]
                last -= 1
                i -= 1
            else:
                i -= 1
        return last + 1


if __name__ == '__main__':
    s = Solution()
    l = [3]
    length = s.removeElement(l, 2)
    print(length, l[:length])
