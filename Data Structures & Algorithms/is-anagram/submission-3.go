func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var count [26]int
	
	for i := range s {
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
