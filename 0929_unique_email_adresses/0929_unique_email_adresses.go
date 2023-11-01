package leetcode

import "strings"

// Time complexity: O(n) where n is strings from emails
// Space complexity: O(n) where n is hashmap length
func numUniqueEmails(emails []string) int {
	hashmap := make(map[string]struct{})

	for _, email := range emails {
			split := strings.Split(email, "@")
			name := strings.Replace(split[0], ".", "", -1)
			name = strings.Split(name, "+")[0] + "@" + split[1]

			hashmap[name] = struct{}{}
	}

	return len(hashmap)
}
