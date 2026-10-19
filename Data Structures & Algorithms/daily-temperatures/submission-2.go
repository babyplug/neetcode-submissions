func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := []int{}

	for i, t := range temperatures {
		for len(stack) > 0 {
			n := len(stack) - 1
			topIndex := stack[n]
			
			if t > temperatures[topIndex] {
				res[topIndex] = i - topIndex
				stack = stack[:n]
			} else {
				break
			}
		}

		stack = append(stack, i)
	}

	return res
}
