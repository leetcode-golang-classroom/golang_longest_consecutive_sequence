package sol

func longestConsecutive(nums []int) int {
	hash := make(map[int]struct{})
	unique := []int{}
	maxLen := 0
	// accumulate
	for _, num := range nums {
		if _, ok := hash[num]; !ok {
			hash[num] = struct{}{}
			unique = append(unique, num)
		}
	}
	// find
	for _, num := range unique {
		if _, ok := hash[num-1]; !ok { // num is smallest
			currentNum := num
			currentLen := 1
			for _, exist := hash[currentNum+1]; exist; _, exist = hash[currentNum+1] {
				currentNum += 1
				currentLen += 1
			}
			if maxLen < currentLen {
				maxLen = currentLen
			}
		}
	}
	return maxLen
}
