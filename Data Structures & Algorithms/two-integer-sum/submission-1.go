func twoSum(nums []int, target int) []int {
	substrMap := make(map[int]int) // key: target - value; value: index
	
	for i := 0; i < len(nums); i++ {
		if _, ok := substrMap[nums[i]]; ok {
			return []int{substrMap[nums[i]], i}
		}
		substrMap[target - nums[i]] = i
	}

	return []int{0, 0}
}
