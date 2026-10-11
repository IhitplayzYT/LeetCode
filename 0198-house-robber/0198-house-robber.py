class Solution:
    def rob(self, nums: list[int]) -> int:
        l = len(nums)
        dp = [-1] * l
        def recurse(i):
            if i >= l:
                return 0
            if dp[i] != -1:
                return dp[i]
            dp[i] = max(nums[i] + recurse(i+2),recurse(i+1))
            return dp[i]

        return max(recurse(0),recurse(1))

