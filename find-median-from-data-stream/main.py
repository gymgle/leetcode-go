import heapq


class MedianFinder(object):

    def __init__(self):
        """
        initialize your data structure here.
        """
        self.minHeap = []
        self.maxHeap = []

    def addNum(self, num):
        """
        :type num: int
        :rtype: None
        """

        if len(self.minHeap) == len(self.maxHeap):
            heapq.heappush(self.minHeap, num)
            heapq.heappush(self.maxHeap, -heapq.heappop(self.minHeap))
        else:
            heapq.heappush(self.maxHeap, -num)
            heapq.heappush(self.minHeap, -heapq.heappop(self.maxHeap))

    def findMedian(self):
        """
        :rtype: float
        """
        if len(self.minHeap) == len(self.maxHeap):
            return float(self.minHeap[0] - self.maxHeap[0]) / 2
        else:
            return -self.maxHeap[0]


# Your MedianFinder object will be instantiated and called as such:
# obj = MedianFinder()
# obj.addNum(num)
# param_2 = obj.findMedian()
