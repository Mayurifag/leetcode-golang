package leetcode

func topKFrequent(nums []int, k int) []int {
	countMap := map[int]int{}
	countSlice := make([][]int, len(nums)+1)

	for _, n := range nums {
		countMap[n]++
	}

	for num, count := range countMap {
		countSlice[count] = append(countSlice[count], num)
	}

	res := make([]int, 0, k)

	for i := len(countSlice) - 1; i > 0; i-- {
    if countSlice[i] != nil {
      res = append(res, countSlice[i]...)
      if len(res) == k {
        break
      }
    }
	}

	return res
}
