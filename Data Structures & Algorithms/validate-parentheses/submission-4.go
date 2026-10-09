

func isValid(s string) bool {
	if len(s) % 2 != 0 {
		return false
	}
    stack := make([]rune, 0, len(s))

	for _, val := range s {
		if val == '(' || val == '{' || val == '['{
            stack = append(stack, val)
        } else {
			if len(stack) == 0 {
				return false
			}
			pop := stack[len(stack)-1]
			if val == ')' && pop != '(' {
				return false
			} else if val == '}' && pop != '{' {
				return false
			} else if val == ']' && pop != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		}

	}
	
	return len(stack) == 0
}
