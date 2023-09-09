package count_paris

func countKDifference(nums []int, k int) int {
	res := 0
	n := len(nums)
	i, j := 0, n-1
	for i < j {
		if nums[i]-nums[j] == k || nums[j]-nums[i] == k {
			res++
		}
		j--
		if i == j {
			i++
			j = n - 1
		}
	}
	return res
}

/* Another solution
// func countKDifference(nums []int, k int) int {
//     numCount := make(map[int]int)
//     res := 0

//     for _, num := range nums {
//         // Check if num + k exists in the set
//         res += numCount[num+k]
//         // Check if num - k exists in the set
//         res += numCount[num-k]
//         // Add the current num to the set
//         numCount[num]++
//     }

//     return res
// }
*/
