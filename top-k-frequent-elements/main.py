class Solution(object):
    def topKFrequent(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: List[int]
        """
        m = dict()
        rst = list()
        for i in nums:
            m.setdefault(i, 0)
            m[i] += 1
        v = sorted(m.items(), key=lambda x:x[1], reverse=True)
        for i in v[:k]:
            rst.append(i[0])
        return rst

if __name__ == '__main__':
    s = Solution()
    print(s.topKFrequent([1, 1, 1, 2, 2, 3], 2))
