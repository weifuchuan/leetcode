package leetcode

import "testing"

func TestFirstMissingPositive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name:"1", args:args{nums:[]int{2,3,4,6}}, want:1},
		{name:"2", args:args{nums:[]int{1,2,0}}, want:3},
		{name:"3", args:args{nums:[]int{3,4,-1,1}}, want:2},
		{name:"4", args:args{nums:[]int{7,8,9,11,12}}, want:1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstMissingPositive(tt.args.nums); got != tt.want {
				t.Errorf("FirstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
