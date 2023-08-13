func containsDuplicate(nums []int) bool {
	mem := map[int]bool{}

	for _, n := range nums {
		if _, exists := mem[n]; exists {
			return true
		} else {
			mem[n] = true
		}
	}
	return false
}