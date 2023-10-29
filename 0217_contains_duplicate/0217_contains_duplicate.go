package leetcode

// Time complexity: O(n) because we loop through each element of nums once
// Space complexity: O(n) because we made hashmap with each element as key in worst case
func containsDuplicate(nums []int) bool {
	if len(nums) == 1 {
		return false
	}

	hashmap := make(map[int]int)

	for _, value := range nums {
		if hashmap[value] == 1 {
			return true
		}
		hashmap[value] = 1
	}

	return false
}


// Proposed solution, better things are:
// struct{}{} takes less memory than int or bool (0 bytes instead of boolean 1)
// shortcut if _, hasNum := hashmap[value]; hasNum {}
//
// func containsDuplicate(nums []int) bool {
// 	hashmap := make(map[int]struct{})

// 	for _, value := range nums {
//	if _, hasNum := hashmap[value]; hasNum {
//		return true
//	}
// 		hashmap[value] = struct{}{}
// 	}

// 	return false
// }
