package alg

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	max := nums[0]
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] >= 0 {
			dp[i] = dp[i-1] + nums[i]

		} else {
			dp[i] = nums[i]
		}
		if dp[i] > max {
			max = dp[i]
		}
	}

	return max
}
