func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	topK := make([]int, 0, k)

	// O(n)
	for _, n := range nums {
		freqMap[n]++
	}

	// O(n*k) (worst case - O(n^2))
	for elem, n := range freqMap {
		if len(topK) < k {
			topK = append(topK, elem)
			continue
		}

		minIndx := 0
		for i := 1; i < k; i++ {
			if freqMap[topK[i]] < freqMap[topK[minIndx]] {
				minIndx = i
			}
		}

		if n > freqMap[topK[minIndx]] {
			topK[minIndx] = elem
		}
	}

	return topK
}

// nums = [1,2,2,3,3,3], k = 2