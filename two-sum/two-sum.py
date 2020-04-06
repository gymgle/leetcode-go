# -*- code: utf-8 -*-

class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        m = {}
        for i in range(0, len(nums)):
            tmp = m.get(target-nums[i], '')
            if  tmp != '':
                return [tmp, i]
            m[nums[i]] = i
        return []
