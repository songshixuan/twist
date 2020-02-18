package alg

func maxProduct(nums []int) int {
	if len(nums) < 2 {
		return nums[0]
	}

	dpMax := make([]int, len(nums))
	dpMin := make([]int, len(nums))

	dpMax[0], dpMin[0] = nums[0], nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			dpMax[i-1], dpMin[i-1] = dpMin[i-1], dpMax[i-1]
		}
		maxTmp := nums[i] * dpMax[i-1]
		if maxTmp > nums[i] {
			dpMax[i] = maxTmp
		} else {
			dpMax[i] = nums[i]
		}
		minTmp := nums[i] * dpMin[i-1]
		if minTmp < nums[i] {
			dpMin[i] = minTmp
		} else {
			dpMin[i] = nums[i]
		}
		if max < dpMax[i] {
			max = dpMax[i]
		}
	}
	return max
}
