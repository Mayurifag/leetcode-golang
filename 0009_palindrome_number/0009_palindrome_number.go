package leetcode

// Type: divmod
func isPalindrome(x int) bool {
	// negative always false
	if x < 0 {
		return false
	}

	number := x
	reverse := 0

	// Time complexity: loop on entire number digit by digit - O(log_10(n))
	// Space complexity: 2 vars used so O(1)
	for number > 0 {
		reverse = reverse * 10 + number % 10
		number /= 10
	}

	return x == reverse
}
