func largestRectangleArea(heights []int) int {
	type Bar struct {
		index int
		height int
	}

	stack := make([]Bar, len(heights))
	maxArea := 0
	for i, h := range heights {
		start := i

		for h < stack[len(stack)-1].height && len(stack) > 0 {
			top := stack[len(stack)-1]
			// Pop top stack
			stack = stack[:len(stack)-1]

			start = top.index

			width := i - top.index
			area := width * top.height

			maxArea = max(area, maxArea)
		}

		stack = append(stack, Bar{
			index: start,
			height: h,
		})
	}

	n := len(heights)

	for _, s := range stack {
		width := n - s.index
		area := width * s.height

		maxArea = max(maxArea, area)
	}

	return maxArea
}
