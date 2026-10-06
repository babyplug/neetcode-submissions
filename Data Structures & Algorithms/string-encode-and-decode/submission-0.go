type Solution struct{}

// Encode is serialize a string to be 1 string separte by 
// length of string and #
func (s *Solution) Encode(strs []string) string {
	var encoded string

	for _, str := range strs {
		encoded += fmt.Sprintf("%d#%s", len(str), str)
	}

	return encoded
}

// Decode is deserialize a string from encoded string 
// back to it's original
func (s *Solution) Decode(encoded string) []string {
	res := make([]string, 0)
	i := 0

	for i < len(encoded) {
		// Find until we found the #
		j := i
		
		// Walks until we found '#'
		for encoded[j] != '#' {
			j++
		}

		// substring and parse it to int
		length, _ := strconv.Atoi(encoded[i:j])
		start := j + 1
		end := start + length
		str := encoded[start:end]
		res = append(res, str)
		i = end
	}

	return res
}
