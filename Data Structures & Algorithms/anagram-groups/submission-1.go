// n - length of strs
// l - length of each word

// time complexity: O(n*l) + O(n) = O(n*l)
func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[[26]int][]string)
	
	// O(n*l)
	for _, word := range strs {
		var charCount [26]int
		for j := 0; j < len(word); j++ {
			charCount[word[j] - 'a']++
		}
		anagramMap[charCount] = append(anagramMap[charCount], word)
	}

	// O(n)
	result := make([][]string, 0, len(anagramMap))
	for _, set := range anagramMap {
		result = append(result, set)
	}

	return result
}

// strs = ["act","pots","tops","cat","stop","hat"]