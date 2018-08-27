package leetcode

/*
Given an unsorted integer array, find the smallest missing positive integer.

Example 1:

Input: [1,2,0]
Output: 3
Example 2:

Input: [3,4,-1,1]
Output: 2
Example 3:

Input: [7,8,9,11,12]
Output: 1
Note:

Your algorithm should run in O(n) time and uses constant extra space.
*/
func FirstMissingPositive(nums []int) int {
	if len(nums)==0{
		return 1
	}
	min:=nums[0]
	vm:=make(map[int]struct{})
	for _, n:=range nums{
		if min>n{
			min=n 
		}
		vm[n]=struct{}{}
	}
	if min>1{
		return 1 
	}
	if min<1{
		min=1
	}
	for i:=min;  ; i++{
		if _,ok:=vm[i];!ok{
			return i 
		}
	}
}