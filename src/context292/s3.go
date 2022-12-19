package context292

import "strconv"

func countTexts(pressedKeys string) int {
	nums := []int{3, 3, 3, 3, 3, 4, 3, 3}
	dp := make([]int, len(pressedKeys))
	dp[0] = 1

	for i := 1; i < len(pressedKeys); i++ {
		curNum, _ := strconv.Atoi(string(pressedKeys[i]))
		num := nums[curNum]
		for j := 1; j <= num; j++ {
			if i-j >= 0 && pressedKeys[i-j] == pressedKeys[i] {
				dp[i] = dp[i] + dp[i-j] + 1
			} else {
				break
			}
		}

	}

	return dp[len(pressedKeys)-1]

}
