func isAnagram(s string, t string) bool {
	var characters [26]int

	if len(s) != len(t) {
		return false
	}

	for i :=0; i<len(s); i++ {
		characters[s[i] - 'a']++
		characters[t[i] - 'a']--
	}

	for _, n := range characters {
		if n != 0 {
			return false
		}
	}

	return true
}
