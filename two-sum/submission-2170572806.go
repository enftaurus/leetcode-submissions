func twoSum(nums []int, target int) []int {
	st := make(map[int]int)

	for idx, num := range nums {
		if idx2, exists := st[target-num]; exists {
			return []int{idx, idx2}
		}
		st[num] = idx
	}

	return nil
}