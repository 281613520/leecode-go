package main

func PancakeSort(arr []int) (ans []int) {
	for i := len(arr); i > 1; i-- {
		index := 0
		for j, v := range arr[:i] {
			if arr[index] < v {
				index = j
			}
		}

		if index == i-1 {
			continue
		}

		for j, m := 0, index; j < (m+1)/2; j++ {
			arr[j], arr[m-j] = arr[m-j], arr[j]
		}
		for j := 0; j < i/2; j++ {
			arr[j], arr[i-1-j] = arr[i-1-j], arr[j]
		}
		ans = append(ans, index+1, i)
	}
	return ans
}
