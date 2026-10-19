func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := []int{}

	for i, t := range temperatures {
		for len(stack) > 0 {
			n := len(stack) - 1
			top := stack[n]
			
			if t > temperatures[top] {
				res[top] = i - top
				stack = stack[:n]
			} else {
				break
			}
		}

		stack = append(stack, i)
	}

	return res
}
