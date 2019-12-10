package permutation

// Ints returns a slice contaning all possible permutations of a slice of ints.
func Ints(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	} else if len(nums) == 1 {
		return [][]int{nums}
	}

	var results [][]int
	for _, sp := range Ints(nums[1:]) {
		results = append(results, allInserts(nums[0], sp)...)
	}

	return results
}

func allInserts(n int, nums []int) [][]int {
	nbResults := len(nums) + 1
	results := make([][]int, nbResults)
	for index := 0; index < nbResults; index++ {
		result := make([]int, nbResults)

		for i := 0; i < index; i++ {
			result[i] = nums[i]
		}
		result[index] = n
		for i := index + 1; i < nbResults; i++ {
			result[i] = nums[i-1]
		}

		results[index] = result
	}

	return results
}
