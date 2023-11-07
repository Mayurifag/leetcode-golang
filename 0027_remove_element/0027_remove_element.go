package leetcode

// Space: O(1) (in place change)
// Time: O(n) (iteration through array)
func removeElement(nums []int, val int) int {
	i := 0
	for _, v := range nums {
		if v != val {
			nums[i] = v
			i++
		}
	}
	return i
}
