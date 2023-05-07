package main

func longestArithSeqLength(nums []int) int {
	minv, maxv := nums[0], nums[0]
	for _, num := range nums[1:] {
		if num < minv {
			minv = num
		} else if num > maxv {
			maxv = num
		}
	}
	diff := maxv - minv
	ans := 1
	for d := -diff; d <= diff; d++ {
		f := make([]int, maxv+1)
		for _, num := range nums {
			prev := num - d
			if prev >= minv && prev <= maxv {
				f[num] = max(f[num], f[prev]+1)
				ans = max(ans, f[num])
			}
			f[num] = max(f[num], 1)
		}
	}
	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
