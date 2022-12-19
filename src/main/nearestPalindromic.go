package main

import (
	"math"
	"strconv"
)

func nearestPalindromic(n string) string {
	// 分4种情况考虑
	// 1. 直接用前面的转
	// 2. 防止过大 需要前面 -1 再反转
	// 3.防止小 前面+1 再反转
	// 4. 100...001 和 99...99

	m := len(n)

	candicates := make([]int, 0)
	candicates = append(candicates, int(math.Pow10(m-1))-1) // 99...99
	candicates = append(candicates, int(math.Pow10(m))+1)   // 100 .. 001

	prefix, _ := strconv.Atoi(n[:(m+1)/2])

	for _, x := range []int{prefix - 1, prefix, prefix + 1} {
		y := x
		if m&1 == 1 {
			y /= 10
		}
		for ; y > 0; y /= 10 {
			x = x*10 + y%10
		}
		candicates = append(candicates, x) //进行三种情况的倒装
	}

	ans := -1
	selfNumber, _ := strconv.Atoi(n)
	for _, candidate := range candicates {
		if candidate != selfNumber {
			if ans == -1 ||
				abs(candidate-selfNumber) < abs(ans-selfNumber) ||
				abs(candidate-selfNumber) == abs(ans-selfNumber) && candidate < ans {
				ans = candidate
			}
		}
	}
	return strconv.Itoa(ans)
}
