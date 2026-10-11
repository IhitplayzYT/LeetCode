class Solution:
    def minCostClimbingStairs(self, cost: list[int]) -> int:
        l = len(cost)
        dp = [-1] * l
        def recurse(i):
            if i < 0:
                return 0
            if i == 0 or i == 1:
                return cost[i]
            if dp[i] == -1:
                dp[i] = cost[i] + min(recurse(i-1),recurse(i-2))
            return dp[i]

        return min(recurse(l-1),recurse(l-2))