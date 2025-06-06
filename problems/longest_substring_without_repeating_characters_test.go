package problems

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	cases := []struct {
		s        string
		expected int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{"au", 2},
		{"aab", 2},
		{"tmmzuxt", 5},
	}

	for _, testCase := range cases {
		t.Run(fmt.Sprintf("%+v", testCase), func(tt *testing.T) {
			assert.Equal(t, testCase.expected, LengthOfLongestSubstring(testCase.s), testCase.s)
		})
	}
}
