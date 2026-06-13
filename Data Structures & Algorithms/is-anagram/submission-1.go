func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := make(map[rune]int)

	for _, value := range s {
		count[value]++
	}

	for _, value := range t {
		count[value]--
		
		if count[value] < 0 {
			return false
		}
	}

	return true
}
