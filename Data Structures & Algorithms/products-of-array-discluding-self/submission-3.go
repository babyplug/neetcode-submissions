func productExceptSelf(nums []int) []int {
	// nums=[1,2,4,6]
	// n = 4
	n := len(nums)
	result := make([]int, n)

	// prefix=[0,0,0,0]
	prefix := make([]int, n)
	// prefix=[1,0,0,0]
	prefix[0] = 1
	// prefix=[1, 1, 2, 8]
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}

	// suffix=[0,0,0,0]
	suffix := make([]int, n)
	// suffix=[0,0,0,1]
	suffix[n-1] = 1
	// nums=[1,2,4,6]
	// i = 2 -> suffix=[0, 0, 6, ...]
	// i = 1 -> suffix=[0, 24, ...]
	// i = 0 -> suffix=[48, ...]
	for i := n-2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i+1]
	}

	for i := range nums {
		result[i] = prefix[i] * suffix[i]
	}

	return result
}
