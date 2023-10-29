package leetcode

// twoSum via bruteforce, -- O(n²) time complexity, O(1) space complexity
func twoSum(nums []int, target int) []int {
	// loop through nums = O(n-1) time complexity
	for i := 0; i < len(nums)-1; i++ {
		// loop through nums = O(n-1) time complexity
		for j := i + 1; j < len(nums); j++ {
			// check if nums[i] + nums[j] == target = O(1) space complexity
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
