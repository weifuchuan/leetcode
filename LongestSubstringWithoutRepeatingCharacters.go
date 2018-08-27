package leetcode

/*
Given a string, find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3 
Explanation: The answer is "abc", which the length is 3.
Example 2:

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3. 
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/
func LengthOfLongestSubstring(s string) int {
	max := 0
	for i := 0; i < len(s); i++ {
		length := 0
		alpSet := make(map[byte]struct{})
		for j := i; j < len(s); j++ {
			if _, ok := alpSet[s[j]]; ok {
				break
			} else {
				alpSet[s[j]] = struct{}{}
				length++
			}
		}
		if max < length {
			max = length
		}
	}
	return max
}
