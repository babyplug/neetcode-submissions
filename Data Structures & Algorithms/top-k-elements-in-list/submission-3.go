
// nums := []int{7, 7, 7, 8, 8, 9}
func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int) // number/frequent

	for _, num := range nums {
		freq[num]++
	}

	// len = 6
	// If we create = 6 -> [0, 1, 2, 3, 4, 5]
	// If we create = 6 + 1 -> [0, 1, 2, 3, 4, 5, 6]
	bucket := make([][]int, len(nums)+1)
	for num, count := range freq {
		bucket[count] = append(bucket[count], num)
	}

	result := make([]int, 0, k)
	for count := len(bucket) - 1; count >= 0; count-- {
		for _, num := range bucket[count] {
			result = append(result, num)

			if len(result) == k {
				return result
			}
		}
	}

    return result
}
