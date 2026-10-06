type Solution struct{}

// Encode is serialize a string to be 1 string separte by 
// length of string and #
func (s *Solution) Encode(strs []string) string {
	var sb strings.Builder

	for _, str := range strs {
		sb.WriteString(strconv.Itoa(len(str)))
		sb.WriteByte('#')
		sb.WriteString(str)
	}

	return sb.String()
}

// Decode is deserialize a string from encoded string 
// back to it's original
func (s *Solution) Decode(encoded string) []string {
	if encoded == "" {
		return nil
	}

	res := make([]string, 0)

	for i := 0; i < len(encoded); {
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
		res = append(res, encoded[start:end])
		i = end
	}

	return res
}
