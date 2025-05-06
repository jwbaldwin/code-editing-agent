# code-editing-agent

## AGENT BUILT:Longest Substring Without Repeating Characters

This repository includes a solution to the classic LeetCode problem "Longest Substring Without Repeating Characters." It was generated entirely by the code-editing-agent.
So was this readme. Except that part, and this.

### Problem Description

Given a string `s`, find the length of the longest substring without repeating characters.

**Examples:**

- `"abcabcbb"` → 3 (The answer is "abc")
- `"bbbbb"` → 1 (The answer is "b")
- `"pwwkew"` → 3 (The answer is "wke")

### Usage

```go
import "agent"

// Find the length of the longest substring without repeating characters
length := agent.LengthOfLongestSubstring("abcabcbb")
fmt.Println(length) // Output: 3
```

### Algorithm

The solution uses a sliding window approach with a hash map:
1. Keep track of the last position of each character
2. When encountering a repeating character, move the window's starting position
3. Calculate the current substring length and update the maximum if needed

Time Complexity: O(n) where n is the length of the string
Space Complexity: O(min(m, n)) where m is the size of the character set
