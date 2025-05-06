package agent

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{"au", 2},
		{"dvdf", 3},
		{"anviaj", 5},
		{"aab", 2},
	}

	for _, tc := range testCases {
		result := LengthOfLongestSubstring(tc.input)
		if result != tc.expected {
			t.Errorf("LengthOfLongestSubstring(%q) = %d; want %d", tc.input, result, tc.expected)
		}
	}
}
