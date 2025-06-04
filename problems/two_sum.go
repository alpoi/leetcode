package problems

// Given an array of integers `nums` and an integer `target`,
// return indices of the two numbers such that they add up to `target`.
// You may assume that each input would have exactly one solution,
// and you may not use the same element twice.

func TwoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)
	dupes := make(map[int]struct{})
	for i, num := range nums {
		if _, exists := dupes[num]; exists {
			continue
		}

		if _, exists := indexMap[num]; exists {
			// If we get a duplicate, it will either be half of the target,
			// or it won't be part of a valid solution, as we are guaranteed
			// exactly one solution.
			if target-num == num {
				return []int{i, indexMap[num]}
			}

			delete(indexMap, num)
			dupes[num] = struct{}{}
			continue
		}

		indexMap[num] = i
	}

	for num, i := range indexMap {
		diff := target - num
		if j, exists := indexMap[diff]; exists && i != j {
			return []int{i, j}
		}
	}

	return nil
}

func TwoSumImproved(nums []int, target int) []int {
	indexMap := make(map[int]int)
	for i, num := range nums {
		diff := target - num
		if j, exists := indexMap[diff]; exists {
			return []int{i, j}
		}
		indexMap[num] = i
	}

	return nil
}
