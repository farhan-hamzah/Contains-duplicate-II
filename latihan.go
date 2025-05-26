package main

import (
	"fmt"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	indexMap := make(map[int]int)

	for i, num := range nums {
		if lastIndex, found := indexMap[num]; found {
			if i - lastIndex <= k {
				return true
			}
		}
		indexMap[num] = i
	}
	return false
}

func main() {
	// Contoh penggunaan
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))  // true
	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))  // true
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 4}, 2))  // false
}
