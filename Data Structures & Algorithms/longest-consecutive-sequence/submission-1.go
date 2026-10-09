func longestConsecutive(nums []int) int {
	sets := make(map[int]struct{})
	
	for _, num := range nums {
		sets[num] = struct{}{}
	}

	res := 0
	for num := range sets {
		_, found := sets[num-1]
		if found {
			continue
		}
		
		current := num
		for {
			_, found := sets[current+1]
			if !found {
				break
			}
			current++
		}
		res = max(res, current - num + 1)
	}

	return res
}
