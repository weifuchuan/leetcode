package leetcode

/*
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.

Example 1:

nums1 = [1, 3]
nums2 = [2]

The median is 2.0
Example 2:

nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5
*/
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var i, j int = 0, 0
	l := 0
	values := make(map[int]int)
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			values[l] = nums1[i]
			i++
		} else {
			values[l] = nums2[j]
			j++
		}
		l++
	}
	for k := i; k < len(nums1); k++ {
		values[l] = nums1[k]
		l++
	}
	for k := j; k < len(nums2); k++ {
		values[l] = nums2[k]
		l++
	}
	var mid float64
	if l%2 == 0 {
		mid = float64(values[l/2-1]+values[l/2]) / 2.0
	} else {
		mid = float64(values[l/2])
	}
	return mid
}
