func isAnagram(s string, t string) bool {
	if len(s) == 0 || len(s) != len(t) {
		return false
	}

	count := map[byte]int{}
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}

	for _, n := range count {
		if n != 0 {
			return false
		}
	}

	return true
}
