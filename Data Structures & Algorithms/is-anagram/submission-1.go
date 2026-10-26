func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charCounter := make(map[rune]int, len(s))

	for _, char := range(s) {
		charCounter[char] ++
	}

	for _, char := range(t) {
		if _, ok := charCounter[char]; !ok {
			return false
		}
		charCounter[char] --
		if charCounter[char] == 0 {
			delete(charCounter, char)
		}
	}

	if len(charCounter) == 0 {
		return true
	} else {
		return false
	}
}
