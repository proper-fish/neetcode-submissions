func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charMap := make(map[byte]int, len(s))

	for i := 0; i < len(s); i++ {
		charMap[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		charMap[t[i]]--
		if charMap[t[i]] < 0 {
			return false
		}
	}
	
	return true
}
