class Solution(object):
    def maxProfit(self, prices):
        """
        :type prices: List[int]
        :rtype: int
        """
        if not prices:
            return 0

        min_price = prices[0]
        max_profit = 0
        for p in prices:
            min_price = min(min_price, p)
            max_profit = max(max_profit, p - min_price)

        return max_profit


if __name__ == '__main__':
    prices = [7, 1, 5, 3, 6, 4]

    s = Solution()
    print(s.maxProfit(prices))
