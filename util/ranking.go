package util

import "sort"

// Rank will rank an int array by giving the rank in a length-matching int array of each corresponding index having the rank
// where the larget int is the highest rank and vise versa
func Rank(scores []int, isReversed bool) []int {
	// Create a copy of the scores slice, so we don't modify the original
	scoresCopy := make([]int, len(scores))
	copy(scoresCopy, scores)

	// Create a mapping of scores to ranks
	scoreMap := make(map[int]int, len(scores))
	sort.Ints(scoresCopy)
	if !isReversed {
		reverse(scoresCopy)
	}

	for i := range scoresCopy {
		if i != 0 && scoresCopy[i] == scoresCopy[i-1] {
			scoreMap[scoresCopy[i]] = scoreMap[scoresCopy[i-1]]
		} else {
			scoreMap[scoresCopy[i]] = i + 1
		}
	}

	// Use the mapping to assign ranks to the original scores
	ranks := make([]int, len(scores))
	for i := range scores {
		ranks[i] = scoreMap[scores[i]]
	}
	return ranks
}

func reverse(a []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
}
