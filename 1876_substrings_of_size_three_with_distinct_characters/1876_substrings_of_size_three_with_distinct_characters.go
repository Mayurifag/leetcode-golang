package leetcode

import "strings"

func countGoodSubstrings(s string) int {
	length := 3
	leftIdx := 0
	hashmap := make(map[string]byte)
	result := 0
	str := strings.Split(s, "")

	for rightIdx, current := range str {
			hashmap[current] += 1

			if rightIdx - leftIdx == length {
					leftChar := str[leftIdx]
					hashmap[leftChar] -= 1

					if hashmap[leftChar] == 0 {
							delete(hashmap, leftChar)
					}

					leftIdx += 1
			}

			if len(hashmap) == length {
					result += 1
			}
	}

	return result
}
