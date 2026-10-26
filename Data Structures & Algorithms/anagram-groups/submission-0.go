// n - length of strs
// l - length of each word

// time complexity: O(n*l) + O(n) = O(n*l)
func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[[26]int][]int)
	
	// O(n*l)
	for i, word := range strs {
		var charCount [26]int
		for j := 0; j < len(word); j++ {
			charCount[word[j] - 'a']++
		}
		anagramMap[charCount] = append(anagramMap[charCount], i)
	}

	// O(n)
	var result [][]string
	for _, set := range anagramMap {  // a. max len(n), b. len = 1
		var setSlice []string
		for _, index := range set { // a. len = 1, b. max len(n)
			setSlice = append(setSlice, strs[index])
		}
		result = append(result, setSlice)
	}

	return result
}

// strs = ["act","pots","tops","cat","stop","hat"]