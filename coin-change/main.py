class Solution(object):
    def coinChange(self, coins, amount):
        """
        :type coins: List[int]
        :type amount: int
        :rtype: int
        """
        def dp(target):
            if target in memo:
                return memo[target]

            if target == 0:
                return 0
            elif target < 0:
                return -1

            mini = int(1e9)
            for c in coins:
                rst = dp(target - c)
                if 0 <= rst < mini:
                    mini = rst + 1
            memo[target] = mini if mini < int(1e9) else -1
            return memo[target]

        memo = dict()
        return dp(amount)


if __name__ == '__main__':
    coins = [1, 2, 5]
    amount = 11

    s = Solution()
    print s.coinChange(coins, amount)

