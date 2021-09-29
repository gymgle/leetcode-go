class Solution(object):
    def __init__(self):
        self.result = 0

    def findTargetSumWays(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """
        def back_track(index, sum):
            if index == len(nums):
                if sum == target:
                    self.result += 1
                return

            back_track(index + 1, sum + nums[index])
            back_track(index + 1, sum - nums[index])

        back_track(0, 0)
        return self.result


if __name__ == '__main__':
    nums = [11, 31, 37, 36, 43, 40, 50, 18, 10, 15, 10, 35, 43, 25, 41, 43, 6, 22, 38, 38]
    target = 44
    s = Solution()
    print s.findTargetSumWays(nums, target)

