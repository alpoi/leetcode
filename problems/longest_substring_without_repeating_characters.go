package problems

import "fmt"

// Given a string `s`, find the length of the longest substring without
// duplicate characters

func LengthOfLongestSubstring(s string) int {
	fmt.Println(s)
	if len(s) == 0 {
		return 0
	}

	// Two pointers, both starting at index 0.
	// Increment the right pointer, recording each character we see.
	// If we encounter a character we've seen before, decrement the right pointer.
	// Then, increment the left pointer and remove it from the seen record.
	// Each time we increment the right pointer, save the substring length.
	// Continue with these rules until the right pointer reaches the end.

	l, r := 0, 0
	seen := map[byte]struct{}{s[0]: {}}
	max := 0

	for {
		r++

		diff := r - l
		if diff > max {
			max = diff
		}

		if r == len(s) {
			break
		}

		if _, ok := seen[s[r]]; ok {
			r--
			delete(seen, s[l])
			l++
			continue
		}

		seen[s[r]] = struct{}{}
	}

	return max
}
