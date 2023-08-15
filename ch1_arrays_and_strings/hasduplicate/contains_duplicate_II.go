package hasduplicate

func ContainsNearbyDuplicate(nums []int, k int) bool {
	elementIndexMap := make(map[int]int)
	for i, num := range nums {
		if j, ok := elementIndexMap[num]; ok {
			if i-j <= k {
				return true
			}
		}
		elementIndexMap[num] = i
	}
	return false
}
