func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charCounter := make(map[rune]int, len(s))

	for _, char := range(s) {
		charCounter[char] ++
	}

	for _, char := range(t) {
		charCounter[char] --
		if charCounter[char] < 0 {
			return false
		}
	}
	return true

	// if len(charCounter) == 0 {
	// 	return true
	// } else {
	// 	return false
	// }
}
