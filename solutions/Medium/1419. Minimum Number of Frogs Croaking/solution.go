package main

func minNumberOfFrogs(croakOfFrogs string) int {
	mp := map[rune]int{'c': 0, 'r': 1, 'o': 2, 'a': 3, 'k': 4}
	cnt := [5]int{}
	frogNum := 0
	res := 0
	for _, ch := range croakOfFrogs {
		t := mp[ch]
		if t == 0 {
			cnt[t]++
			frogNum++
			if frogNum > res {
				res = frogNum
			}
		} else {
			if cnt[t-1] == 0 {
				return -1
			}
			cnt[t-1]--
			if t == 4 {
				frogNum--
			} else {
				cnt[t]++
			}
		}
	}
	if frogNum > 0 {
		return -1
	}
	return res
}
