func topKFrequent(nums []int, k int) []int {
	// key - element, value - frequency
	freqMap := make(map[int]int)
	
	// index - frequency, value - elements of such frequency
	buckets := make([][]int, len(nums) + 1)

	// O(n)
	for _, n := range nums {
		freqMap[n]++
	}

	for elem, n := range freqMap {
		buckets[n] = append(buckets[n], elem)
	}

	topK := make([]int, 0, k)
	for i := len(nums); i > 0; i-- {
		if len(buckets[i]) > 0 {
			topK = append(topK, buckets[i]...)
		}
		if len(topK) >= k {
			break
		}
	}

	return topK
}

// nums = [1,2,2,3,3,3], k = 2