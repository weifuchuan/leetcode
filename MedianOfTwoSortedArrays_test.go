package leetcode

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "1", args: args{nums1: []int{1, 2, 3, 4}, nums2: []int{2, 3, 4, 5}}, want: 3.0},
		{name: "2", args: args{nums1: []int{1, 3}, nums2: []int{2}}, want: 2.0},
		{name: "3", args: args{nums1: []int{1, 2}, nums2: []int{3, 4}}, want: 2.5},
		{name: "4", args: args{nums1: []int{5,6,7,10}, nums2: []int{3, 4,6,11}}, want: 6},
		{name: "5", args: args{nums1: []int{9, 10, 19}, nums2: []int{11,15,16,19}}, want: 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("FindMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
