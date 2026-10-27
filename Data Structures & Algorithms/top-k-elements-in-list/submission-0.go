func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	topK := make([]int, 0, k)

	for _, n := range nums {
		freqMap[n]++
	}

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