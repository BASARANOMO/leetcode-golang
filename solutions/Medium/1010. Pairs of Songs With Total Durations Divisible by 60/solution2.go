package main

func numPairsDivisibleBy60(time []int) int {
	cnt := [60]int{}
	var res int
	for _, t := range time {
		res += cnt[(60-t%60)%60]
		cnt[t%60] += 1
	}
	return res
}
