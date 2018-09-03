package leetcode

import "testing"

func TestIsMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
	}{
		{name:"1", args:args{s:"", p:""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			IsMatch(tt.args.s, tt.args.p)
		})
	}
}
