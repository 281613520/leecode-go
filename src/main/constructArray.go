package main

func constructArray(n int, k int) []int {
	ans := make([]int, 0)

	// 1,n,2,n-1,3,n-2,n-3....4
	// 1,n,2,n-1,3,4,...n-2

	l := 1
	r := n
	flag := true

	for i := k; i > 0; i-- {
		if flag {
			ans = append(ans, l)
			l++
		} else {
			ans = append(ans, r)
			r--
		}
		flag = !flag
	}
	size := len(ans)

	if flag {
		for i := n - size; i > 0; i-- {
			ans = append(ans, r)
			r--
		}
	} else {
		for i := n - size; i > 0; i-- {
			ans = append(ans, l)
			l++
		}
	}

	return ans
}
