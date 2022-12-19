package tcontest80

func countSubarrays(nums []int, k int64) int64 {
	l := 0
	ans := int64(0)
	sum := int64(0)

	for r, num := range nums {
		sum += int64(num)
		for sum*int64(r-l+1) >= k {
			sum -= int64(nums[l])
			l++
		}

		ans += int64(r - l + 1)
	}

	return ans
}
