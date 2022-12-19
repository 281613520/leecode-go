package main

func countQuadruplets(nums []int) int {
	ans := 0

	n := len(nums)
	for l := 3; l < n; l++ {
		for k := 2; k < l; k++ {
			for j := 1; j < k; j++ {
				for i := 0; i < j; i++ {
					if nums[i]+nums[j]+nums[k] == nums[l] {
						ans++
					}
				}
			}
		}
	}

	return ans
}
