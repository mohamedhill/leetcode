func maxOperations(nums []int, k int) int {
	count := make(map[int]int)
	res := 0

	for _, num := range nums {
		need := k - num

		if count[need] > 0 {
			res++
			count[need]--
		} else {
			count[num]++
		}
	}

	return res
}