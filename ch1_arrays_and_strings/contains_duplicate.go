package main

import "fmt"

func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)

	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 3, 3, 5, 6, 7}))
	fmt.Println(containsDuplicate([]int{1, 3, 4, 5, 6, 7}))
}
