package agent

// LengthOfLongestSubstring solves the "Longest Substring Without Repeating Characters" problem
// Given a string s, it finds the length of the longest substring without repeating characters.
//
// Problem description:
// - Input: A string s
// - Output: The length of the longest substring without repeating characters
//
// Example 1:
// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
//
// Example 2:
// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
//
// Example 3:
// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
func LengthOfLongestSubstring(s string) int {
	// Map to store the last seen position of each character
	charMap := make(map[byte]int)

	// Variables to track the maximum length and the current starting position
	maxLength := 0
	start := 0

	for i := range s {
		// If the character has been seen before and is in the current substring
		if pos, exists := charMap[s[i]]; exists && pos >= start {
			// Move the start pointer to the position after the last occurrence
			start = pos + 1
		}

		// Update the last seen position of the current character
		charMap[s[i]] = i

		// Update the maximum length if needed
		maxLength = max(maxLength, i-start+1)
	}

	return maxLength
}
