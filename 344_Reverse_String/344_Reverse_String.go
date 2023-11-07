package leetcode

// Space - O(1)
// Time - O(n/2) => O(n)
func reverseString(s []byte)  {
	leftIdx := 0
	rightIdx := len(s)-1

	for leftIdx < rightIdx {
		s[leftIdx], s[rightIdx] = s[rightIdx], s[leftIdx]
		leftIdx++
		rightIdx--
	}
}
