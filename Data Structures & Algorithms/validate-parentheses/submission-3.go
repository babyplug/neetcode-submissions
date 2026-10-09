

func isValid(s string) bool {
	if len(s) % 2 != 0 {
		return false
	}
    stack := make([]byte, 0, len(s))

	for i := range s {
		switch s[i] {
		case '(', '{', '[':
			stack = append(stack, s[i])
		default:
			if len(stack) == 0 {
				return false
			}
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if s[i] == ')' && pop != '(' {
				return false
			} else if s[i] == '}' && pop != '{' {
				return false
			} else if s[i] == ']' && pop != '[' {
				return false
			}
		}
	}
	
	return len(stack) == 0
}
