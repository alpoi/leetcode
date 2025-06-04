package problems

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	cases := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{0, 4, 3, 0}, 0, []int{0, 3}},
	}

	for _, testCase := range cases {
		t.Run(fmt.Sprintf("%+v", testCase), func(tt *testing.T) {
			assert.ElementsMatch(tt, testCase.expected, TwoSum(testCase.nums, testCase.target))
		})
	}
}

func TestTwoSumImproved(t *testing.T) {
	cases := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{0, 4, 3, 0}, 0, []int{0, 3}},
	}
	for _, testCase := range cases {
		t.Run(fmt.Sprintf("%+v", testCase), func(tt *testing.T) {
			assert.ElementsMatch(tt, testCase.expected, TwoSumImproved(testCase.nums, testCase.target))
		})
	}
}
